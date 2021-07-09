package terraform

import (
	"fmt"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/jckuester/awstools-lib/internal"
	"github.com/jckuester/awstools-lib/terraform/provider"
	"github.com/jckuester/terradozer/pkg/resource"
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
	resource.UpdatableResource
}

// ResourcesThreadSafe is a list implementation to store resources concurrently.
type ResourcesThreadSafe struct {
	sync.Mutex
	Resources []Resource
	Errors    []error
}

// UpdateStates updates the Terraform state for each resource via the Terraform AWS Provider.
// Returns only resources for which an update was successful. Errors about failed updates are returned via []error.
// If the existingOnly flag is true, only existing resources are returned
// (i.e. which state isn't of type cty.Nil after update).
func UpdateStates(resources []Resource, providerPath string, parallel int, existingOnly bool) ([]Resource, []error) {
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

			p, err := provider.Launch(providerPath, 1*time.Minute)
			if err != nil {
				result.Lock()
				result.Errors = append(result.Errors,
					fmt.Errorf("failed to launch provider (%s): %s", providerPath, err))
				result.Unlock()

				return
			}
			defer p.Close()

			config := cty.ObjectVal(map[string]cty.Value{
				"profile":                     cty.StringVal(r.Profile),
				"region":                      cty.StringVal(r.Region),
				"access_key":                  cty.UnknownVal(cty.DynamicPseudoType),
				"allowed_account_ids":         cty.UnknownVal(cty.DynamicPseudoType),
				"assume_role":                 cty.UnknownVal(cty.DynamicPseudoType),
				"default_tags":                cty.UnknownVal(cty.DynamicPseudoType),
				"endpoints":                   cty.UnknownVal(cty.DynamicPseudoType),
				"forbidden_account_ids":       cty.UnknownVal(cty.DynamicPseudoType),
				"ignore_tag_prefixes":         cty.UnknownVal(cty.DynamicPseudoType),
				"ignore_tags":                 cty.UnknownVal(cty.DynamicPseudoType),
				"insecure":                    cty.UnknownVal(cty.DynamicPseudoType),
				"max_retries":                 cty.UnknownVal(cty.DynamicPseudoType),
				"s3_force_path_style":         cty.UnknownVal(cty.DynamicPseudoType),
				"secret_key":                  cty.UnknownVal(cty.DynamicPseudoType),
				"shared_credentials_file":     cty.UnknownVal(cty.DynamicPseudoType),
				"skip_credentials_validation": cty.UnknownVal(cty.DynamicPseudoType),
				"skip_get_ec2_platforms":      cty.UnknownVal(cty.DynamicPseudoType),
				"skip_metadata_api_check":     cty.UnknownVal(cty.DynamicPseudoType),
				"skip_region_validation":      cty.UnknownVal(cty.DynamicPseudoType),
				"skip_requesting_account_id":  cty.UnknownVal(cty.DynamicPseudoType),
				"token":                       cty.UnknownVal(cty.DynamicPseudoType),
			})

			err = p.Configure(config)
			if err != nil {
				result.Lock()
				result.Errors = append(result.Errors, fmt.Errorf("failed to configure provider: %s", err))
				result.Unlock()

				return
			}

			r.UpdatableResource = resource.New(r.Type, r.ID, nil, p)

			log.WithFields(log.Fields{
				"type":    r.Type,
				"id":      r.ID,
				"region":  r.Region,
				"profile": r.Profile,
			}).Debugf("start updating Terraform state of resource")

			err = r.UpdateState()
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

			if r.State() == nil {
				return
			}

			// filter out resources that don't exist anymore
			// (e.g., ECS clusters in state INACTIVE)
			if existingOnly && r.State().IsNull() {
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
