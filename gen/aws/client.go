//go:build codegen

package aws

import (
	"bytes"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

// GenerateClient returns Go code that initializes all AWS API clients that are used by Terraform resources.
func GenerateClient(outputPath string, services map[string]string) {
	var servicesUsedByTerraform []string
	for _, service := range services {
		if _, ok := excludeServices[service]; ok {
			continue
		}

		serviceV2, ok := AWSServicesV1toV2[service]
		if ok {
			service = serviceV2
		}

		if !util.Contains(servicesUsedByTerraform, service) {
			servicesUsedByTerraform = append(servicesUsedByTerraform, service)
		}
	}

	servicesUsedByTerraform = append(servicesUsedByTerraform, "sts")
	sort.Strings(servicesUsedByTerraform)

	err := util.WriteGoFile(
		filepath.Join(outputPath, "client.go"),
		util.CodeLayout,
		"",
		"aws",
		clientGoCode(servicesUsedByTerraform),
	)

	if err != nil {
		panic(err)
	}
}

func clientGoCode(services []string) string {
	var buf bytes.Buffer
	err := clientTmpl.Execute(&buf, services)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var clientTmpl = template.Must(template.New("client").Funcs(
	template.FuncMap{
		"Title": strings.Title,
	}).Parse(`import (
"context"
"fmt"
"time"
"github.com/apex/log"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/sts"
{{range .}}"github.com/aws/aws-sdk-go-v2/service/{{.}}"
{{end -}}
"github.com/pkg/errors"
)

type Client struct {
	AccountID string
	Region string
	Profile string
	{{ range . }}{{ . | Title }}conn *{{.}}.Client
{{end}}}

func NewClient(ctx context.Context, configs ...func(*config.LoadOptions) error) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
	{{ range . }}{{ . | Title }}conn: {{.}}.NewFromConfig(cfg),
	{{end}}}

	client.Region = cfg.Region

	log.WithFields(log.Fields{
		"region": cfg.Region,
	}).Debugf("created new instance of AWS client")

	return client, nil
}

// SetAccountID populates the AccountID field of the client.
func (client *Client) SetAccountID(ctx context.Context) error {
	resp, err := client.Stsconn.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return errors.Wrap(err, "failed to get caller identity")
	}

	client.AccountID = *resp.Account

	return nil
}
`))
