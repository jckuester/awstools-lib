// +build codegen

package aws

import (
	"sort"
)

type Service struct {
	Name                   string
	TerraformResourceTypes []ResourceType
}

type ResourceType struct {
	Name         string
	Tags         bool
	CreationTime bool
	Owner        bool
}

var excludeServices = map[string]struct{}{
	"importexport":      {},
	"prometheusservice": {},
	"simpledb":          {},
}

var AWSServicesV1toV2 = map[string]string{
	"elb":   "elasticloadbalancing",
	"elbv2": "elasticloadbalancingv2",
}

// ResourceTypesForAWSServices returns the Terraform resource types belonging to each AWS service.
// An AWS service is only part of the result if its associated resource type list is not empty.
func ResourceTypesForAWSServices(serviceMap map[string]string) []Service {
	rTypesByService := make(map[string][]ResourceType)

	for rType, service := range serviceMap {
		rTypes, ok := rTypesByService[service]
		if !ok {
			rTypesByService[service] = []ResourceType{{Name: rType}}
			continue
		}
		rTypesByService[service] = append(rTypes, ResourceType{Name: rType})
	}

	var result []Service
	for service, rTypes := range rTypesByService {
		sort.Slice(rTypes, func(i, j int) bool {
			return rTypes[i].Name < rTypes[j].Name
		})

		result = append(result, Service{
			Name:                   service,
			TerraformResourceTypes: rTypes,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}
