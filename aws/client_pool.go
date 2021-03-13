package aws

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
)

// clientPoolThreadSafe is a concurrent map implementation to store multiple AWS clients.
type clientPoolThreadSafe struct {
	sync.Mutex
	clients map[ClientKey]Client
}

type ClientKey struct {
	Profile, Region string
}

// NewClientPool creates an AWS client for each permutation of the given profiles and regions.
// If profiles, regions, or both are empty, credentials and regions are picked up via the usual default provider chain,
// respectively. For example, if regions are empty, the region is first looked for via the according region environment variable
// or second the default region for each profile is used from `~/.aws/config`.
func NewClientPool(ctx context.Context, profiles []string, regions []string) (map[ClientKey]Client, error) {
	errors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	clientPool := &clientPoolThreadSafe{
		clients: make(map[ClientKey]Client),
	}

	if len(profiles) > 0 && len(regions) > 0 {
		wg.Add(len(profiles) * len(regions))

		for _, profile := range profiles {
			for _, region := range regions {

				go func(p string, r string) {
					defer wg.Done()

					client, err := NewClient(
						ctx,
						config.WithSharedConfigProfile(p),
						config.WithRegion(r))
					if err != nil {
						errors <- err
						return
					}

					clientPool.Lock()
					clientPool.clients[ClientKey{p, client.Region}] = *client
					clientPool.Unlock()
				}(profile, region)
			}
		}
	} else if len(profiles) > 0 {
		wg.Add(len(profiles))

		for _, profile := range profiles {
			go func(p string) {
				defer wg.Done()

				client, err := NewClient(ctx, config.WithSharedConfigProfile(p))
				if err != nil {
					errors <- err
					return
				}

				clientPool.Lock()
				clientPool.clients[ClientKey{p, client.Region}] = *client
				clientPool.Unlock()
			}(profile)
		}
	} else if len(regions) > 0 {
		wg.Add(len(regions))

		for _, region := range regions {
			go func(r string) {
				defer wg.Done()

				client, err := NewClient(ctx, config.WithRegion(r))
				if err != nil {
					errors <- err
					return
				}

				clientPool.Lock()
				clientPool.clients[ClientKey{"", client.Region}] = *client
				clientPool.Unlock()
			}(region)
		}
	} else {
		client, err := NewClient(ctx)
		if err != nil {
			return nil, err
		}

		return map[ClientKey]Client{{"", client.Region}: *client}, nil
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errors:
		close(errors)
		return nil, err
	}

	return clientPool.clients, nil
}
