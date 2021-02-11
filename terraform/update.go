package terraform

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/apex/log"

	"github.com/fatih/color"
	awsls "github.com/jckuester/awsls/aws"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/internal"
	"github.com/jckuester/terradozer/pkg/provider"
	terradozerRes "github.com/jckuester/terradozer/pkg/resource"
)

// ResourcesThreadSafe is a list implementation to store resources concurrently.
type ResourcesThreadSafe struct {
	sync.Mutex
	Resources []awsls.Resource
}

// UpdateStates fetches the Terraform state for each resource via the Terraform AWS Provider.
// Returns only resources which still exist (i.e. state isn't of type cty.Nil after update).
func UpdateStates(resources []awsls.Resource, providers map[aws.ClientKey]provider.TerraformProvider) []awsls.Resource {
	var wg sync.WaitGroup

	result := &ResourcesThreadSafe{
		Resources: []awsls.Resource{},
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

			log.WithFields(log.Fields{
				"type":    r.Type,
				"id":      r.ID,
				"region":  r.Region,
				"profile": r.Profile,
				"time":    time.Now().Format("04:05"),
			}).Debugf("start updating Terraform state of resource")

			err := r.UpdateState()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
				return
			}

			log.WithFields(log.Fields{
				"type":    r.Type,
				"id":      r.ID,
				"region":  r.Region,
				"profile": r.Profile,
				"time":    time.Now().Format("04:05"),
			}).Debugf("updated Terraform state of resource")

			if r.State() == nil {
				return
			}

			// filter out resources that don't exist anymore
			// (e.g., ECS clusters in state INACTIVE)
			if r.State().IsNull() {
				return
			}

			result.Lock()
			result.Resources = append(result.Resources, *r)
			result.Unlock()
		}(i)
	}

	// Wait for all updates to complete
	wg.Wait()

	return resources
}
