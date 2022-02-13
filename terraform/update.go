package terraform

import (
	"fmt"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/internal"
	"github.com/jckuester/awstools-lib/terraform/provider"
	"github.com/zclconf/go-cty/cty"
)

type Resource struct {
	// Type is a resource's type as defined by the Terraform AWS Provider (e.g., aws_instance).
	Type string
	// ID is a resource's ID as defined by the Terraform AWS Provider. Depending on the resource type, the ID can be
	// the ARN, name (for aws_iam_role), or combined string of different attributes. The Terraform ID can differ from
	// what AWS defines as ID and from what ID is returned via the AWS SDK.
	ID string
	// The AWS region where a resource lives in.
	Region    string
	Profile   string
	AccountID string
	Tags      map[string]string
	CreatedAt *time.Time

	// Provider is the Terraform Provider to update the state of a resource
	Provider *provider.TerraformProvider
	// State is the Terraform state of the resource
	State *cty.Value
	Attrs map[string]cty.Value
}

// ResourcesThreadSafe is a list implementation to store resources concurrently.
type ResourcesThreadSafe struct {
	sync.Mutex
	Resources []Resource
	Errors    []error
}

// UpdatableResource implementations can update a Terraform resource's state.
type UpdatableResource interface {
	Type() string
	ID() string
	State() *cty.Value
	UpdateState() error
}

// UpdateStates updates the Terraform state for each resource via the Terraform AWS Provider.
// Returns only resources for which an update was successful. Errors about failed updates are returned via []error.
// If the existingOnly flag is true, only existing resources are returned
// (i.e. which state isn't of type cty.Nil after update).
func UpdateStates(resources []Resource, providers map[aws.ClientKey]provider.TerraformProvider,
	parallel int, existingOnly bool) ([]Resource, []error) {
	var wg sync.WaitGroup

	result := &ResourcesThreadSafe{
		Resources: []Resource{},
		Errors:    []error{},
	}

	sem := internal.NewSemaphore(parallel)
	for i := range resources {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Acquire a semaphore so that we can limit concurrency
			sem.Acquire()
			defer sem.Release()

			r := &resources[i]

			key := aws.ClientKey{
				Profile: r.Profile,
				Region:  r.Region,
			}

			p, ok := providers[key]

			if !ok {
				panic(fmt.Sprintf("could not find Terraform AWS Provider for key: %v", key))
			}

			r.Provider = &p

			log.WithFields(log.Fields{
				"type":    r.Type,
				"id":      r.ID,
				"region":  r.Region,
				"profile": r.Profile,
			}).Debugf("start updating Terraform state of resource")

			err := r.UpdateState()
			if err != nil {
				result.Lock()
				result.Errors = append(result.Errors, err)
				result.Unlock()

				return
			}

			log.WithFields(log.Fields{
				"type":    r.Type,
				"id":      r.ID,
				"region":  r.Region,
				"profile": r.Profile,
			}).Debugf("updated Terraform state of resource")

			if r.State == nil {
				return
			}

			// filter out resources that don't exist anymore
			// (e.g., ECS clusters in state INACTIVE)
			if existingOnly && r.State.IsNull() {
				return
			}

			result.Lock()
			result.Resources = append(result.Resources, *r)
			result.Unlock()
		}(i)
	}

	// Wait for all updates to complete
	wg.Wait()

	return result.Resources, result.Errors
}

// UpdateResources updates the state of a given list of resources in parallel.
// Only updated resources are returned which still exist in AWS.
func UpdateResources(resources []UpdatableResource, parallel int) []UpdatableResource {
	numOfResourcesToUpdate := len(resources)

	var updatedResources []UpdatableResource

	jobQueue := make(chan UpdatableResource, numOfResourcesToUpdate)

	workerResults := make(chan updateWorkerResult, numOfResourcesToUpdate)

	for workerID := 1; workerID <= parallel; workerID++ {
		go updateWorker(jobQueue, workerResults)
	}

	for _, r := range resources {
		jobQueue <- r
	}

	close(jobQueue)

	for i := 1; i <= numOfResourcesToUpdate; i++ {
		r := <-workerResults

		if r.err != nil {
			log.WithError(r.err).WithFields(log.Fields{
				"type":        r.resource.Type(),
				"resource_id": r.resource.ID(),
			}).Info("cannot refresh resource state")

			continue
		}

		updatedResources = append(updatedResources, r.resource)
	}

	return updatedResources
}

type updateWorkerResult struct {
	resource UpdatableResource
	// err is set if update failed.
	err error
}

// updateWorker is a worker that updates the state of a resource.
func updateWorker(resources <-chan UpdatableResource, result chan<- updateWorkerResult) {
	for r := range resources {
		err := r.UpdateState()
		if err != nil {
			result <- updateWorkerResult{resource: r, err: err}

			continue
		}

		resourceNotFound := r.State().IsNull()
		if resourceNotFound {
			result <- updateWorkerResult{resource: r, err: fmt.Errorf("resource doesn't exist anymore")}

			continue
		}

		result <- updateWorkerResult{resource: r, err: nil}
	}
}

// UpdateState updates the state of the resource (i.e., refreshes all its attributes).
// If the resource is already gone, the updated state will be nil (more precisely, of type cty.NilVal).
func (r *Resource) UpdateState() error {
	if r.State != nil {
		// if the resource stores already a state representation, refresh that state
		result, err := r.Provider.ReadResource(r.Type, *r.State)
		if err != nil {
			return fmt.Errorf("failed to read current state of resource: %s", err)
		}

		r.State = &result

		return nil
	}

	result, err := r.importAndReadResource()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"id": r.ID, "type": r.Type}).
			Debug("failed to import resource; trying to read resource without import")

		result, err = r.readResource()
		if err != nil {
			return err
		}
	}

	r.State = &result

	return nil
}

func (r Resource) importAndReadResource() (cty.Value, error) {
	importedResources, err := r.Provider.ImportResource(r.Type, r.ID)
	if err != nil {
		return cty.NilVal, err
	}

	for _, rImported := range importedResources {
		currentResourceState, err := r.Provider.ReadResource(rImported.TypeName, rImported.State)
		if err != nil {
			return cty.NilVal, err
		}

		if rImported.TypeName == r.Type {
			return currentResourceState, nil
		}

		log.WithError(err).WithFields(log.Fields{
			"type": rImported.TypeName,
		}).Debug("found multiple resources during import")
	}

	return cty.NilVal, fmt.Errorf("no resource found to be imported")
}

// readResource fetches the current state of a resource based on its ID attribute.
func (r Resource) readResource() (cty.Value, error) {
	schema, err := r.Provider.GetSchemaForResource(r.Type)
	if err != nil {
		return cty.NilVal, err
	}

	if r.Attrs == nil {
		r.Attrs = map[string]cty.Value{}
	}

	r.Attrs["id"] = cty.StringVal(r.ID)

	currentResourceState, err := r.Provider.ReadResource(r.Type, emptyValueWithAttrs(r.Attrs, schema.Block))
	if err != nil {
		return cty.NilVal, fmt.Errorf("failed to read current state of resource: %s", err)
	}

	return currentResourceState, nil
}

// emptyValueWithAttrs returns a non-null object for the configuration block
// where all attribute values are set to empty values except the given ones.
//
// see also github.com/hashicorp/terraform/configs/configschema/empty_value.go
func emptyValueWithAttrs(attrs map[string]cty.Value, block *configschema.Block) cty.Value {
	vals := make(map[string]cty.Value)

	for name, attrS := range block.Attributes {
		attr, ok := attrs[name]
		if ok {
			vals[name] = attr
		} else {
			vals[name] = attrS.EmptyValue()
		}
	}

	for name, blockS := range block.BlockTypes {
		vals[name] = blockS.EmptyValue()
	}

	return cty.ObjectVal(vals)
}
