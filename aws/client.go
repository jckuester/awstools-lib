// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

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
	"github.com/pkg/errors"
)

type Client struct {
	AccountID                     string
	Region                        string
	Profile                       string
	Accessanalyzerconn            *accessanalyzer.Client
	Acmconn                       *acm.Client
	Acmpcaconn                    *acmpca.Client
	Apigatewayconn                *apigateway.Client
	Apigatewayv2conn              *apigatewayv2.Client
	Applicationautoscalingconn    *applicationautoscaling.Client
	Appmeshconn                   *appmesh.Client
	Appsyncconn                   *appsync.Client
	Athenaconn                    *athena.Client
	Autoscalingconn               *autoscaling.Client
	Autoscalingplansconn          *autoscalingplans.Client
	Backupconn                    *backup.Client
	Batchconn                     *batch.Client
	Budgetsconn                   *budgets.Client
	Cloud9conn                    *cloud9.Client
	Cloudformationconn            *cloudformation.Client
	Cloudfrontconn                *cloudfront.Client
	Cloudhsmv2conn                *cloudhsmv2.Client
	Cloudtrailconn                *cloudtrail.Client
	Cloudwatchconn                *cloudwatch.Client
	Cloudwatcheventsconn          *cloudwatchevents.Client
	Cloudwatchlogsconn            *cloudwatchlogs.Client
	Codeartifactconn              *codeartifact.Client
	Codebuildconn                 *codebuild.Client
	Codecommitconn                *codecommit.Client
	Codedeployconn                *codedeploy.Client
	Codepipelineconn              *codepipeline.Client
	Codestarconnectionsconn       *codestarconnections.Client
	Codestarnotificationsconn     *codestarnotifications.Client
	Cognitoidentityconn           *cognitoidentity.Client
	Cognitoidentityproviderconn   *cognitoidentityprovider.Client
	Configserviceconn             *configservice.Client
	Costandusagereportserviceconn *costandusagereportservice.Client
	Databasemigrationserviceconn  *databasemigrationservice.Client
	Datapipelineconn              *datapipeline.Client
	Datasyncconn                  *datasync.Client
	Daxconn                       *dax.Client
	Devicefarmconn                *devicefarm.Client
	Directconnectconn             *directconnect.Client
	Directoryserviceconn          *directoryservice.Client
	Dlmconn                       *dlm.Client
	Docdbconn                     *docdb.Client
	Dynamodbconn                  *dynamodb.Client
	Ec2conn                       *ec2.Client
	Ecrconn                       *ecr.Client
	Ecrpublicconn                 *ecrpublic.Client
	Ecsconn                       *ecs.Client
	Efsconn                       *efs.Client
	Eksconn                       *eks.Client
	Elasticacheconn               *elasticache.Client
	Elasticbeanstalkconn          *elasticbeanstalk.Client
	Elasticloadbalancingconn      *elasticloadbalancing.Client
	Elasticloadbalancingv2conn    *elasticloadbalancingv2.Client
	Elasticsearchserviceconn      *elasticsearchservice.Client
	Elastictranscoderconn         *elastictranscoder.Client
	Emrconn                       *emr.Client
	Firehoseconn                  *firehose.Client
	Fmsconn                       *fms.Client
	Fsxconn                       *fsx.Client
	Gameliftconn                  *gamelift.Client
	Glacierconn                   *glacier.Client
	Globalacceleratorconn         *globalaccelerator.Client
	Glueconn                      *glue.Client
	Guarddutyconn                 *guardduty.Client
	Iamconn                       *iam.Client
	Imagebuilderconn              *imagebuilder.Client
	Inspectorconn                 *inspector.Client
	Iotconn                       *iot.Client
	Kafkaconn                     *kafka.Client
	Kinesisconn                   *kinesis.Client
	Kinesisanalyticsconn          *kinesisanalytics.Client
	Kinesisanalyticsv2conn        *kinesisanalyticsv2.Client
	Kinesisvideoconn              *kinesisvideo.Client
	Kmsconn                       *kms.Client
	Lakeformationconn             *lakeformation.Client
	Lambdaconn                    *lambda.Client
	Lexmodelbuildingserviceconn   *lexmodelbuildingservice.Client
	Licensemanagerconn            *licensemanager.Client
	Lightsailconn                 *lightsail.Client
	Macieconn                     *macie.Client
	Mediaconvertconn              *mediaconvert.Client
	Mediapackageconn              *mediapackage.Client
	Mediastoreconn                *mediastore.Client
	Mqconn                        *mq.Client
	Neptuneconn                   *neptune.Client
	Networkfirewallconn           *networkfirewall.Client
	Opsworksconn                  *opsworks.Client
	Organizationsconn             *organizations.Client
	Pinpointconn                  *pinpoint.Client
	Qldbconn                      *qldb.Client
	Quicksightconn                *quicksight.Client
	Ramconn                       *ram.Client
	Rdsconn                       *rds.Client
	Redshiftconn                  *redshift.Client
	Resourcegroupsconn            *resourcegroups.Client
	Route53conn                   *route53.Client
	Route53resolverconn           *route53resolver.Client
	S3conn                        *s3.Client
	S3controlconn                 *s3control.Client
	S3outpostsconn                *s3outposts.Client
	Sagemakerconn                 *sagemaker.Client
	Secretsmanagerconn            *secretsmanager.Client
	Securityhubconn               *securityhub.Client
	Servicecatalogconn            *servicecatalog.Client
	Servicediscoveryconn          *servicediscovery.Client
	Servicequotasconn             *servicequotas.Client
	Sesconn                       *ses.Client
	Sfnconn                       *sfn.Client
	Shieldconn                    *shield.Client
	Signerconn                    *signer.Client
	Snsconn                       *sns.Client
	Sqsconn                       *sqs.Client
	Ssmconn                       *ssm.Client
	Ssoadminconn                  *ssoadmin.Client
	Storagegatewayconn            *storagegateway.Client
	Stsconn                       *sts.Client
	Swfconn                       *swf.Client
	Syntheticsconn                *synthetics.Client
	Transferconn                  *transfer.Client
	Wafconn                       *waf.Client
	Wafregionalconn               *wafregional.Client
	Wafv2conn                     *wafv2.Client
	Worklinkconn                  *worklink.Client
	Workspacesconn                *workspaces.Client
	Xrayconn                      *xray.Client
}

func NewClient(ctx context.Context, configs ...func(*config.LoadOptions) error) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		Accessanalyzerconn:            accessanalyzer.NewFromConfig(cfg),
		Acmconn:                       acm.NewFromConfig(cfg),
		Acmpcaconn:                    acmpca.NewFromConfig(cfg),
		Apigatewayconn:                apigateway.NewFromConfig(cfg),
		Apigatewayv2conn:              apigatewayv2.NewFromConfig(cfg),
		Applicationautoscalingconn:    applicationautoscaling.NewFromConfig(cfg),
		Appmeshconn:                   appmesh.NewFromConfig(cfg),
		Appsyncconn:                   appsync.NewFromConfig(cfg),
		Athenaconn:                    athena.NewFromConfig(cfg),
		Autoscalingconn:               autoscaling.NewFromConfig(cfg),
		Autoscalingplansconn:          autoscalingplans.NewFromConfig(cfg),
		Backupconn:                    backup.NewFromConfig(cfg),
		Batchconn:                     batch.NewFromConfig(cfg),
		Budgetsconn:                   budgets.NewFromConfig(cfg),
		Cloud9conn:                    cloud9.NewFromConfig(cfg),
		Cloudformationconn:            cloudformation.NewFromConfig(cfg),
		Cloudfrontconn:                cloudfront.NewFromConfig(cfg),
		Cloudhsmv2conn:                cloudhsmv2.NewFromConfig(cfg),
		Cloudtrailconn:                cloudtrail.NewFromConfig(cfg),
		Cloudwatchconn:                cloudwatch.NewFromConfig(cfg),
		Cloudwatcheventsconn:          cloudwatchevents.NewFromConfig(cfg),
		Cloudwatchlogsconn:            cloudwatchlogs.NewFromConfig(cfg),
		Codeartifactconn:              codeartifact.NewFromConfig(cfg),
		Codebuildconn:                 codebuild.NewFromConfig(cfg),
		Codecommitconn:                codecommit.NewFromConfig(cfg),
		Codedeployconn:                codedeploy.NewFromConfig(cfg),
		Codepipelineconn:              codepipeline.NewFromConfig(cfg),
		Codestarconnectionsconn:       codestarconnections.NewFromConfig(cfg),
		Codestarnotificationsconn:     codestarnotifications.NewFromConfig(cfg),
		Cognitoidentityconn:           cognitoidentity.NewFromConfig(cfg),
		Cognitoidentityproviderconn:   cognitoidentityprovider.NewFromConfig(cfg),
		Configserviceconn:             configservice.NewFromConfig(cfg),
		Costandusagereportserviceconn: costandusagereportservice.NewFromConfig(cfg),
		Databasemigrationserviceconn:  databasemigrationservice.NewFromConfig(cfg),
		Datapipelineconn:              datapipeline.NewFromConfig(cfg),
		Datasyncconn:                  datasync.NewFromConfig(cfg),
		Daxconn:                       dax.NewFromConfig(cfg),
		Devicefarmconn:                devicefarm.NewFromConfig(cfg),
		Directconnectconn:             directconnect.NewFromConfig(cfg),
		Directoryserviceconn:          directoryservice.NewFromConfig(cfg),
		Dlmconn:                       dlm.NewFromConfig(cfg),
		Docdbconn:                     docdb.NewFromConfig(cfg),
		Dynamodbconn:                  dynamodb.NewFromConfig(cfg),
		Ec2conn:                       ec2.NewFromConfig(cfg),
		Ecrconn:                       ecr.NewFromConfig(cfg),
		Ecrpublicconn:                 ecrpublic.NewFromConfig(cfg),
		Ecsconn:                       ecs.NewFromConfig(cfg),
		Efsconn:                       efs.NewFromConfig(cfg),
		Eksconn:                       eks.NewFromConfig(cfg),
		Elasticacheconn:               elasticache.NewFromConfig(cfg),
		Elasticbeanstalkconn:          elasticbeanstalk.NewFromConfig(cfg),
		Elasticloadbalancingconn:      elasticloadbalancing.NewFromConfig(cfg),
		Elasticloadbalancingv2conn:    elasticloadbalancingv2.NewFromConfig(cfg),
		Elasticsearchserviceconn:      elasticsearchservice.NewFromConfig(cfg),
		Elastictranscoderconn:         elastictranscoder.NewFromConfig(cfg),
		Emrconn:                       emr.NewFromConfig(cfg),
		Firehoseconn:                  firehose.NewFromConfig(cfg),
		Fmsconn:                       fms.NewFromConfig(cfg),
		Fsxconn:                       fsx.NewFromConfig(cfg),
		Gameliftconn:                  gamelift.NewFromConfig(cfg),
		Glacierconn:                   glacier.NewFromConfig(cfg),
		Globalacceleratorconn:         globalaccelerator.NewFromConfig(cfg),
		Glueconn:                      glue.NewFromConfig(cfg),
		Guarddutyconn:                 guardduty.NewFromConfig(cfg),
		Iamconn:                       iam.NewFromConfig(cfg),
		Imagebuilderconn:              imagebuilder.NewFromConfig(cfg),
		Inspectorconn:                 inspector.NewFromConfig(cfg),
		Iotconn:                       iot.NewFromConfig(cfg),
		Kafkaconn:                     kafka.NewFromConfig(cfg),
		Kinesisconn:                   kinesis.NewFromConfig(cfg),
		Kinesisanalyticsconn:          kinesisanalytics.NewFromConfig(cfg),
		Kinesisanalyticsv2conn:        kinesisanalyticsv2.NewFromConfig(cfg),
		Kinesisvideoconn:              kinesisvideo.NewFromConfig(cfg),
		Kmsconn:                       kms.NewFromConfig(cfg),
		Lakeformationconn:             lakeformation.NewFromConfig(cfg),
		Lambdaconn:                    lambda.NewFromConfig(cfg),
		Lexmodelbuildingserviceconn:   lexmodelbuildingservice.NewFromConfig(cfg),
		Licensemanagerconn:            licensemanager.NewFromConfig(cfg),
		Lightsailconn:                 lightsail.NewFromConfig(cfg),
		Macieconn:                     macie.NewFromConfig(cfg),
		Mediaconvertconn:              mediaconvert.NewFromConfig(cfg),
		Mediapackageconn:              mediapackage.NewFromConfig(cfg),
		Mediastoreconn:                mediastore.NewFromConfig(cfg),
		Mqconn:                        mq.NewFromConfig(cfg),
		Neptuneconn:                   neptune.NewFromConfig(cfg),
		Networkfirewallconn:           networkfirewall.NewFromConfig(cfg),
		Opsworksconn:                  opsworks.NewFromConfig(cfg),
		Organizationsconn:             organizations.NewFromConfig(cfg),
		Pinpointconn:                  pinpoint.NewFromConfig(cfg),
		Qldbconn:                      qldb.NewFromConfig(cfg),
		Quicksightconn:                quicksight.NewFromConfig(cfg),
		Ramconn:                       ram.NewFromConfig(cfg),
		Rdsconn:                       rds.NewFromConfig(cfg),
		Redshiftconn:                  redshift.NewFromConfig(cfg),
		Resourcegroupsconn:            resourcegroups.NewFromConfig(cfg),
		Route53conn:                   route53.NewFromConfig(cfg),
		Route53resolverconn:           route53resolver.NewFromConfig(cfg),
		S3conn:                        s3.NewFromConfig(cfg),
		S3controlconn:                 s3control.NewFromConfig(cfg),
		S3outpostsconn:                s3outposts.NewFromConfig(cfg),
		Sagemakerconn:                 sagemaker.NewFromConfig(cfg),
		Secretsmanagerconn:            secretsmanager.NewFromConfig(cfg),
		Securityhubconn:               securityhub.NewFromConfig(cfg),
		Servicecatalogconn:            servicecatalog.NewFromConfig(cfg),
		Servicediscoveryconn:          servicediscovery.NewFromConfig(cfg),
		Servicequotasconn:             servicequotas.NewFromConfig(cfg),
		Sesconn:                       ses.NewFromConfig(cfg),
		Sfnconn:                       sfn.NewFromConfig(cfg),
		Shieldconn:                    shield.NewFromConfig(cfg),
		Signerconn:                    signer.NewFromConfig(cfg),
		Snsconn:                       sns.NewFromConfig(cfg),
		Sqsconn:                       sqs.NewFromConfig(cfg),
		Ssmconn:                       ssm.NewFromConfig(cfg),
		Ssoadminconn:                  ssoadmin.NewFromConfig(cfg),
		Storagegatewayconn:            storagegateway.NewFromConfig(cfg),
		Stsconn:                       sts.NewFromConfig(cfg),
		Swfconn:                       swf.NewFromConfig(cfg),
		Syntheticsconn:                synthetics.NewFromConfig(cfg),
		Transferconn:                  transfer.NewFromConfig(cfg),
		Wafconn:                       waf.NewFromConfig(cfg),
		Wafregionalconn:               wafregional.NewFromConfig(cfg),
		Wafv2conn:                     wafv2.NewFromConfig(cfg),
		Worklinkconn:                  worklink.NewFromConfig(cfg),
		Workspacesconn:                workspaces.NewFromConfig(cfg),
		Xrayconn:                      xray.NewFromConfig(cfg),
	}

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
