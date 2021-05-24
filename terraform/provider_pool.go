package terraform

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/terradozer/pkg/provider"
	"github.com/zclconf/go-cty/cty"
)

// providerPoolThreadSafe is a concurrent map implementation to store multiple Terraform AWS Providers.
type providerPoolThreadSafe struct {
	sync.Mutex
	providers map[aws.ClientKey]provider.TerraformProvider
}

// NewProviderPool launches a set of Terraform AWS Providers with the configuration of the given clientKeys
// (combination of AWS profile and region).
func NewProviderPool(ctx context.Context, clientKeys []aws.ClientKey, version, installDir string,
	timeout time.Duration) (
	map[aws.ClientKey]provider.TerraformProvider, error) {

	metaPlugin, err := provider.Install("aws", version, installDir)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("failed to install provider (%s): %s", "aws", err))
	}

	errors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	providerPool := &providerPoolThreadSafe{
		providers: make(map[aws.ClientKey]provider.TerraformProvider),
	}

	if len(clientKeys) > 0 {
		wg.Add(len(clientKeys))

		for _, clientKey := range clientKeys {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			go func(p string, r string) {
				defer wg.Done()

				log.WithFields(log.Fields{
					"profile": p,
					"region":  r,
				}).Debugf("start launching new instance of Terraform AWS Provider")

				pr, err := provider.Launch(metaPlugin.Path, timeout)
				if err != nil {
					errors <- fmt.Errorf("failed to launch provider (%s): %s", metaPlugin.Path, err)
					return
				}

				config := cty.ObjectVal(map[string]cty.Value{
					"profile":                     cty.StringVal(p),
					"region":                      cty.StringVal(r),
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

				err = pr.Configure(config)
				if err != nil {
					pr.Close()

					errors <- fmt.Errorf("failed to configure provider (name=%s, version=%s): %s",
						metaPlugin.Name, metaPlugin.Version, err)

					return
				}

				providerPool.Lock()
				providerPool.providers[aws.ClientKey{Profile: p, Region: r}] = *pr
				providerPool.Unlock()

				log.WithFields(log.Fields{
					"profile": p,
					"region":  r,
				}).Debugf("launched new instance of Terraform AWS Provider")
			}(clientKey.Profile, clientKey.Region)
		}
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

		for _, p := range providerPool.providers {
			_ = p.Close()
		}

		return nil, err
	}

	return providerPool.providers, nil
}
