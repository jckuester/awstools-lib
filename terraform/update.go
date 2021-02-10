package terraform

import (
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
	awsls "github.com/jckuester/awsls/aws"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/internal"
	"github.com/jckuester/terradozer/pkg/provider"
	terradozerRes "github.com/jckuester/terradozer/pkg/resource"
)

// resourcesThreadSafe is a list implementation to store resources concurrently.
type resourcesThreadSafe struct {
	sync.Mutex
	resources []awsls.Resource
}

// UpdateStates fetches the Terraform state for each resource via the Terraform AWS Provider.
// Returns only resources which still exist (i.e. state isn't of type cty.Nil after update).
func UpdateStates(resources []awsls.Resource, providers map[aws.ClientKey]provider.TerraformProvider) []awsls.Resource {
	var wg sync.WaitGroup

	result := &resourcesThreadSafe{
		resources: []awsls.Resource{},
	}

	sem := internal.NewSemaphore(5)
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

			r.UpdatableResource = terradozerRes.New(r.Type, r.ID, nil, &p)

			err := r.UpdateState()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
			}

			// filter out resources that don't exist anymore
			// (e.g., ECS clusters in state INACTIVE)
			if r.State() != nil && r.State().IsNull() {
				return
			}

			result.Lock()
			result.resources = append(result.resources, *r)
			result.Unlock()
		}(i)
	}

	// Wait for all updates to complete
	wg.Wait()

	return resources
}
