//go:build codegen
//go:generate go run -tags codegen generate.go
//go:generate gofmt -s -w ../aws ../terraform
//go:generate goimports -w ../aws ../terraform
//go:generate mockgen -source=../terraform/update.go -destination=../terraform/update_mock_test.go -package=terraform_test

package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jckuester/awstools-lib/gen/aws"
	"github.com/jckuester/awstools-lib/gen/terraform"
)

const (
	outputPathAWS                = "../aws"
	outputPathResource           = "../terraform"
	terraformAwsProviderRepoPath = "/home/jan/git/github.com/terraform-provider-aws"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	err := os.MkdirAll(outputPathResource, 0775)
	if err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}

	err = os.MkdirAll(outputPathAWS, 0775)
	if err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}

	resourceTypes, err := terraform.GenerateResourceTypeList(terraformAwsProviderRepoPath, outputPathResource)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list of Terraform AWS resource types")
	}

	resourceFileNames, err := terraform.ResourceFileNames(terraformAwsProviderRepoPath, resourceTypes)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> filename implementing resource")
	}

	serviceByResourceType := terraform.GenerateResourceServiceMap(terraformAwsProviderRepoPath, outputPathResource,
		resourceTypes, resourceFileNames)

	aws.GenerateClient(outputPathAWS, serviceByResourceType)
}
