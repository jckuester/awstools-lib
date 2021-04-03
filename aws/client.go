// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Client struct {
	AccountID                     string
	Region                        string
	Profile                       string
	Iamconn                       *iam.Client
	Apigatewayconn                *apigateway.Client
	Apigatewayv2conn              *apigatewayv2.Client
	Elasticloadbalancingconn      *elasticloadbalancing.Client
	Cloudfrontconn                *cloudfront.Client
	Codecommitconn                *codecommit.Client
	Configserviceconn             *configservice.Client
	Efsconn                       *efs.Client
	Sagemakerconn                 *sagemaker.Client
	Sesconn                       *ses.Client
	Ec2conn                       *ec2.Client
	Wafregionalconn               *wafregional.Client
	Applicationautoscalingconn    *applicationautoscaling.Client
	Codebuildconn                 *codebuild.Client
	Fmsconn                       *fms.Client
	Glacierconn                   *glacier.Client
	S3conn                        *s3.Client
	Lightsailconn                 *lightsail.Client
	Docdbconn                     *docdb.Client
	Glueconn                      *glue.Client
	Opsworksconn                  *opsworks.Client
	Servicecatalogconn            *servicecatalog.Client
	Sqsconn                       *sqs.Client
	Appsyncconn                   *appsync.Client
	Codepipelineconn              *codepipeline.Client
	Elasticacheconn               *elasticache.Client
	Imagebuilderconn              *imagebuilder.Client
	Dlmconn                       *dlm.Client
	Elasticbeanstalkconn          *elasticbeanstalk.Client
	Ecrconn                       *ecr.Client
	Ecsconn                       *ecs.Client
	Lexmodelbuildingserviceconn   *lexmodelbuildingservice.Client
	Autoscalingconn               *autoscaling.Client
	Transferconn                  *transfer.Client
	Accessanalyzerconn            *accessanalyzer.Client
	Securityhubconn               *securityhub.Client
	Iotconn                       *iot.Client
	Route53resolverconn           *route53resolver.Client
	Wafconn                       *waf.Client
	Appmeshconn                   *appmesh.Client
	Storagegatewayconn            *storagegateway.Client
	Shieldconn                    *shield.Client
	Rdsconn                       *rds.Client
	Elasticsearchserviceconn      *elasticsearchservice.Client
	Qldbconn                      *qldb.Client
	Ecrpublicconn                 *ecrpublic.Client
	Databasemigrationserviceconn  *databasemigrationservice.Client
	Route53conn                   *route53.Client
	Emrconn                       *emr.Client
	Athenaconn                    *athena.Client
	Batchconn                     *batch.Client
	Cloudwatchlogsconn            *cloudwatchlogs.Client
	Codedeployconn                *codedeploy.Client
	Guarddutyconn                 *guardduty.Client
	Pinpointconn                  *pinpoint.Client
	Organizationsconn             *organizations.Client
	Directconnectconn             *directconnect.Client
	S3controlconn                 *s3control.Client
	Gameliftconn                  *gamelift.Client
	Ssmconn                       *ssm.Client
	Ssoadminconn                  *ssoadmin.Client
	Wafv2conn                     *wafv2.Client
	Codeartifactconn              *codeartifact.Client
	Elasticloadbalancingv2conn    *elasticloadbalancingv2.Client
	Networkfirewallconn           *networkfirewall.Client
	Cloudwatcheventsconn          *cloudwatchevents.Client
	Redshiftconn                  *redshift.Client
	Signerconn                    *signer.Client
	Acmpcaconn                    *acmpca.Client
	Datapipelineconn              *datapipeline.Client
	Syntheticsconn                *synthetics.Client
	Workspacesconn                *workspaces.Client
	Servicediscoveryconn          *servicediscovery.Client
	Lambdaconn                    *lambda.Client
	Cloudformationconn            *cloudformation.Client
	Lakeformationconn             *lakeformation.Client
	Kafkaconn                     *kafka.Client
	Secretsmanagerconn            *secretsmanager.Client
	Datasyncconn                  *datasync.Client
	Ramconn                       *ram.Client
	Neptuneconn                   *neptune.Client
	Servicequotasconn             *servicequotas.Client
	Costandusagereportserviceconn *costandusagereportservice.Client
	Directoryserviceconn          *directoryservice.Client
	Worklinkconn                  *worklink.Client
	Snsconn                       *sns.Client
	Xrayconn                      *xray.Client
	Mediastoreconn                *mediastore.Client
	Cognitoidentityproviderconn   *cognitoidentityprovider.Client
	Inspectorconn                 *inspector.Client
	Cloud9conn                    *cloud9.Client
	Autoscalingplansconn          *autoscalingplans.Client
	Codestarnotificationsconn     *codestarnotifications.Client
	Mqconn                        *mq.Client
	Mediaconvertconn              *mediaconvert.Client
	Backupconn                    *backup.Client
	Kmsconn                       *kms.Client
	Cognitoidentityconn           *cognitoidentity.Client
	Dynamodbconn                  *dynamodb.Client
	Fsxconn                       *fsx.Client
	Resourcegroupsconn            *resourcegroups.Client
	Budgetsconn                   *budgets.Client
	Codestarconnectionsconn       *codestarconnections.Client
	Daxconn                       *dax.Client
	Firehoseconn                  *firehose.Client
	Macieconn                     *macie.Client
	Mediapackageconn              *mediapackage.Client
	Swfconn                       *swf.Client
	Acmconn                       *acm.Client
	Kinesisanalyticsconn          *kinesisanalytics.Client
	Sfnconn                       *sfn.Client
	Kinesisanalyticsv2conn        *kinesisanalyticsv2.Client
	Eksconn                       *eks.Client
	Kinesisvideoconn              *kinesisvideo.Client
	Globalacceleratorconn         *globalaccelerator.Client
	Cloudwatchconn                *cloudwatch.Client
	Cloudtrailconn                *cloudtrail.Client
	Licensemanagerconn            *licensemanager.Client
	Quicksightconn                *quicksight.Client
	Devicefarmconn                *devicefarm.Client
	Cloudhsmv2conn                *cloudhsmv2.Client
	Elastictranscoderconn         *elastictranscoder.Client
	S3outpostsconn                *s3outposts.Client
	Kinesisconn                   *kinesis.Client
	Stsconn                       *sts.Client
}

func NewClient(ctx context.Context, configs ...func(*config.LoadOptions) error) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		Iamconn:                       iam.NewFromConfig(cfg),
		Apigatewayconn:                apigateway.NewFromConfig(cfg),
		Apigatewayv2conn:              apigatewayv2.NewFromConfig(cfg),
		Elasticloadbalancingconn:      elasticloadbalancing.NewFromConfig(cfg),
		Cloudfrontconn:                cloudfront.NewFromConfig(cfg),
		Codecommitconn:                codecommit.NewFromConfig(cfg),
		Configserviceconn:             configservice.NewFromConfig(cfg),
		Efsconn:                       efs.NewFromConfig(cfg),
		Sagemakerconn:                 sagemaker.NewFromConfig(cfg),
		Sesconn:                       ses.NewFromConfig(cfg),
		Ec2conn:                       ec2.NewFromConfig(cfg),
		Wafregionalconn:               wafregional.NewFromConfig(cfg),
		Applicationautoscalingconn:    applicationautoscaling.NewFromConfig(cfg),
		Codebuildconn:                 codebuild.NewFromConfig(cfg),
		Fmsconn:                       fms.NewFromConfig(cfg),
		Glacierconn:                   glacier.NewFromConfig(cfg),
		S3conn:                        s3.NewFromConfig(cfg),
		Lightsailconn:                 lightsail.NewFromConfig(cfg),
		Docdbconn:                     docdb.NewFromConfig(cfg),
		Glueconn:                      glue.NewFromConfig(cfg),
		Opsworksconn:                  opsworks.NewFromConfig(cfg),
		Servicecatalogconn:            servicecatalog.NewFromConfig(cfg),
		Sqsconn:                       sqs.NewFromConfig(cfg),
		Appsyncconn:                   appsync.NewFromConfig(cfg),
		Codepipelineconn:              codepipeline.NewFromConfig(cfg),
		Elasticacheconn:               elasticache.NewFromConfig(cfg),
		Imagebuilderconn:              imagebuilder.NewFromConfig(cfg),
		Dlmconn:                       dlm.NewFromConfig(cfg),
		Elasticbeanstalkconn:          elasticbeanstalk.NewFromConfig(cfg),
		Ecrconn:                       ecr.NewFromConfig(cfg),
		Ecsconn:                       ecs.NewFromConfig(cfg),
		Lexmodelbuildingserviceconn:   lexmodelbuildingservice.NewFromConfig(cfg),
		Autoscalingconn:               autoscaling.NewFromConfig(cfg),
		Transferconn:                  transfer.NewFromConfig(cfg),
		Accessanalyzerconn:            accessanalyzer.NewFromConfig(cfg),
		Securityhubconn:               securityhub.NewFromConfig(cfg),
		Iotconn:                       iot.NewFromConfig(cfg),
		Route53resolverconn:           route53resolver.NewFromConfig(cfg),
		Wafconn:                       waf.NewFromConfig(cfg),
		Appmeshconn:                   appmesh.NewFromConfig(cfg),
		Storagegatewayconn:            storagegateway.NewFromConfig(cfg),
		Shieldconn:                    shield.NewFromConfig(cfg),
		Rdsconn:                       rds.NewFromConfig(cfg),
		Elasticsearchserviceconn:      elasticsearchservice.NewFromConfig(cfg),
		Qldbconn:                      qldb.NewFromConfig(cfg),
		Ecrpublicconn:                 ecrpublic.NewFromConfig(cfg),
		Databasemigrationserviceconn:  databasemigrationservice.NewFromConfig(cfg),
		Route53conn:                   route53.NewFromConfig(cfg),
		Emrconn:                       emr.NewFromConfig(cfg),
		Athenaconn:                    athena.NewFromConfig(cfg),
		Batchconn:                     batch.NewFromConfig(cfg),
		Cloudwatchlogsconn:            cloudwatchlogs.NewFromConfig(cfg),
		Codedeployconn:                codedeploy.NewFromConfig(cfg),
		Guarddutyconn:                 guardduty.NewFromConfig(cfg),
		Pinpointconn:                  pinpoint.NewFromConfig(cfg),
		Organizationsconn:             organizations.NewFromConfig(cfg),
		Directconnectconn:             directconnect.NewFromConfig(cfg),
		S3controlconn:                 s3control.NewFromConfig(cfg),
		Gameliftconn:                  gamelift.NewFromConfig(cfg),
		Ssmconn:                       ssm.NewFromConfig(cfg),
		Ssoadminconn:                  ssoadmin.NewFromConfig(cfg),
		Wafv2conn:                     wafv2.NewFromConfig(cfg),
		Codeartifactconn:              codeartifact.NewFromConfig(cfg),
		Elasticloadbalancingv2conn:    elasticloadbalancingv2.NewFromConfig(cfg),
		Networkfirewallconn:           networkfirewall.NewFromConfig(cfg),
		Cloudwatcheventsconn:          cloudwatchevents.NewFromConfig(cfg),
		Redshiftconn:                  redshift.NewFromConfig(cfg),
		Signerconn:                    signer.NewFromConfig(cfg),
		Acmpcaconn:                    acmpca.NewFromConfig(cfg),
		Datapipelineconn:              datapipeline.NewFromConfig(cfg),
		Syntheticsconn:                synthetics.NewFromConfig(cfg),
		Workspacesconn:                workspaces.NewFromConfig(cfg),
		Servicediscoveryconn:          servicediscovery.NewFromConfig(cfg),
		Lambdaconn:                    lambda.NewFromConfig(cfg),
		Cloudformationconn:            cloudformation.NewFromConfig(cfg),
		Lakeformationconn:             lakeformation.NewFromConfig(cfg),
		Kafkaconn:                     kafka.NewFromConfig(cfg),
		Secretsmanagerconn:            secretsmanager.NewFromConfig(cfg),
		Datasyncconn:                  datasync.NewFromConfig(cfg),
		Ramconn:                       ram.NewFromConfig(cfg),
		Neptuneconn:                   neptune.NewFromConfig(cfg),
		Servicequotasconn:             servicequotas.NewFromConfig(cfg),
		Costandusagereportserviceconn: costandusagereportservice.NewFromConfig(cfg),
		Directoryserviceconn:          directoryservice.NewFromConfig(cfg),
		Worklinkconn:                  worklink.NewFromConfig(cfg),
		Snsconn:                       sns.NewFromConfig(cfg),
		Xrayconn:                      xray.NewFromConfig(cfg),
		Mediastoreconn:                mediastore.NewFromConfig(cfg),
		Cognitoidentityproviderconn:   cognitoidentityprovider.NewFromConfig(cfg),
		Inspectorconn:                 inspector.NewFromConfig(cfg),
		Cloud9conn:                    cloud9.NewFromConfig(cfg),
		Autoscalingplansconn:          autoscalingplans.NewFromConfig(cfg),
		Codestarnotificationsconn:     codestarnotifications.NewFromConfig(cfg),
		Mqconn:                        mq.NewFromConfig(cfg),
		Mediaconvertconn:              mediaconvert.NewFromConfig(cfg),
		Backupconn:                    backup.NewFromConfig(cfg),
		Kmsconn:                       kms.NewFromConfig(cfg),
		Cognitoidentityconn:           cognitoidentity.NewFromConfig(cfg),
		Dynamodbconn:                  dynamodb.NewFromConfig(cfg),
		Fsxconn:                       fsx.NewFromConfig(cfg),
		Resourcegroupsconn:            resourcegroups.NewFromConfig(cfg),
		Budgetsconn:                   budgets.NewFromConfig(cfg),
		Codestarconnectionsconn:       codestarconnections.NewFromConfig(cfg),
		Daxconn:                       dax.NewFromConfig(cfg),
		Firehoseconn:                  firehose.NewFromConfig(cfg),
		Macieconn:                     macie.NewFromConfig(cfg),
		Mediapackageconn:              mediapackage.NewFromConfig(cfg),
		Swfconn:                       swf.NewFromConfig(cfg),
		Acmconn:                       acm.NewFromConfig(cfg),
		Kinesisanalyticsconn:          kinesisanalytics.NewFromConfig(cfg),
		Sfnconn:                       sfn.NewFromConfig(cfg),
		Kinesisanalyticsv2conn:        kinesisanalyticsv2.NewFromConfig(cfg),
		Eksconn:                       eks.NewFromConfig(cfg),
		Kinesisvideoconn:              kinesisvideo.NewFromConfig(cfg),
		Globalacceleratorconn:         globalaccelerator.NewFromConfig(cfg),
		Cloudwatchconn:                cloudwatch.NewFromConfig(cfg),
		Cloudtrailconn:                cloudtrail.NewFromConfig(cfg),
		Licensemanagerconn:            licensemanager.NewFromConfig(cfg),
		Quicksightconn:                quicksight.NewFromConfig(cfg),
		Devicefarmconn:                devicefarm.NewFromConfig(cfg),
		Cloudhsmv2conn:                cloudhsmv2.NewFromConfig(cfg),
		Elastictranscoderconn:         elastictranscoder.NewFromConfig(cfg),
		S3outpostsconn:                s3outposts.NewFromConfig(cfg),
		Kinesisconn:                   kinesis.NewFromConfig(cfg),
		Stsconn:                       sts.NewFromConfig(cfg),
	}

	client.Region = cfg.Region

	log.WithFields(log.Fields{
		"region": cfg.Region,
		"time":   time.Now().Format("04:05.000"),
	}).Debugf("created new instance of AWS client")

	return client, nil
}

// SetAccountID populates the AccountID field of the client.
func (client *Client) SetAccountID(ctx context.Context) error {
	resp, err := client.Stsconn.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Printf("%T", err.Error())
		return errors.Wrap(err, "failed to get caller identity")
	}

	client.AccountID = *resp.Account

	return nil
}
