// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/braket"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/chime"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/ebs"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/honeycode"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	"github.com/aws/aws-sdk-go-v2/service/iotfleethub"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless"
	"github.com/aws/aws-sdk-go-v2/service/ivs"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go-v2/service/mobile"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/outposts"
	"github.com/aws/aws-sdk-go-v2/service/personalize"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	"github.com/aws/aws-sdk-go-v2/service/pi"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldbsession"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/sms"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Client struct {
	AccountID                           string
	Region                              string
	Profile                             string
	Accessanalyzerconn                  *accessanalyzer.Client
	Acmconn                             *acm.Client
	Acmpcaconn                          *acmpca.Client
	Alexaforbusinessconn                *alexaforbusiness.Client
	Amplifyconn                         *amplify.Client
	Apigatewayconn                      *apigateway.Client
	Apigatewaymanagementapiconn         *apigatewaymanagementapi.Client
	Apigatewayv2conn                    *apigatewayv2.Client
	Appconfigconn                       *appconfig.Client
	Appflowconn                         *appflow.Client
	Applicationautoscalingconn          *applicationautoscaling.Client
	Applicationdiscoveryserviceconn     *applicationdiscoveryservice.Client
	Applicationinsightsconn             *applicationinsights.Client
	Appmeshconn                         *appmesh.Client
	Appstreamconn                       *appstream.Client
	Appsyncconn                         *appsync.Client
	Athenaconn                          *athena.Client
	Auditmanagerconn                    *auditmanager.Client
	Autoscalingconn                     *autoscaling.Client
	Autoscalingplansconn                *autoscalingplans.Client
	Backupconn                          *backup.Client
	Batchconn                           *batch.Client
	Braketconn                          *braket.Client
	Budgetsconn                         *budgets.Client
	Chimeconn                           *chime.Client
	Cloud9conn                          *cloud9.Client
	Clouddirectoryconn                  *clouddirectory.Client
	Cloudformationconn                  *cloudformation.Client
	Cloudfrontconn                      *cloudfront.Client
	Cloudhsmconn                        *cloudhsm.Client
	Cloudhsmv2conn                      *cloudhsmv2.Client
	Cloudsearchconn                     *cloudsearch.Client
	Cloudsearchdomainconn               *cloudsearchdomain.Client
	Cloudtrailconn                      *cloudtrail.Client
	Cloudwatchconn                      *cloudwatch.Client
	Cloudwatcheventsconn                *cloudwatchevents.Client
	Cloudwatchlogsconn                  *cloudwatchlogs.Client
	Codeartifactconn                    *codeartifact.Client
	Codebuildconn                       *codebuild.Client
	Codecommitconn                      *codecommit.Client
	Codedeployconn                      *codedeploy.Client
	Codeguruprofilerconn                *codeguruprofiler.Client
	Codegurureviewerconn                *codegurureviewer.Client
	Codepipelineconn                    *codepipeline.Client
	Codestarconn                        *codestar.Client
	Codestarconnectionsconn             *codestarconnections.Client
	Codestarnotificationsconn           *codestarnotifications.Client
	Cognitoidentityconn                 *cognitoidentity.Client
	Cognitoidentityproviderconn         *cognitoidentityprovider.Client
	Cognitosyncconn                     *cognitosync.Client
	Comprehendconn                      *comprehend.Client
	Comprehendmedicalconn               *comprehendmedical.Client
	Computeoptimizerconn                *computeoptimizer.Client
	Configserviceconn                   *configservice.Client
	Connectconn                         *connect.Client
	Connectcontactlensconn              *connectcontactlens.Client
	Connectparticipantconn              *connectparticipant.Client
	Costandusagereportserviceconn       *costandusagereportservice.Client
	Costexplorerconn                    *costexplorer.Client
	Customerprofilesconn                *customerprofiles.Client
	Databasemigrationserviceconn        *databasemigrationservice.Client
	Dataexchangeconn                    *dataexchange.Client
	Datapipelineconn                    *datapipeline.Client
	Datasyncconn                        *datasync.Client
	Daxconn                             *dax.Client
	Detectiveconn                       *detective.Client
	Devicefarmconn                      *devicefarm.Client
	Devopsguruconn                      *devopsguru.Client
	Directconnectconn                   *directconnect.Client
	Directoryserviceconn                *directoryservice.Client
	Dlmconn                             *dlm.Client
	Docdbconn                           *docdb.Client
	Dynamodbconn                        *dynamodb.Client
	Dynamodbstreamsconn                 *dynamodbstreams.Client
	Ebsconn                             *ebs.Client
	Ec2conn                             *ec2.Client
	Ec2instanceconnectconn              *ec2instanceconnect.Client
	Ecrconn                             *ecr.Client
	Ecrpublicconn                       *ecrpublic.Client
	Ecsconn                             *ecs.Client
	Efsconn                             *efs.Client
	Eksconn                             *eks.Client
	Elasticacheconn                     *elasticache.Client
	Elasticbeanstalkconn                *elasticbeanstalk.Client
	Elasticinferenceconn                *elasticinference.Client
	Elasticloadbalancingconn            *elasticloadbalancing.Client
	Elasticloadbalancingv2conn          *elasticloadbalancingv2.Client
	Elasticsearchserviceconn            *elasticsearchservice.Client
	Elastictranscoderconn               *elastictranscoder.Client
	Emrconn                             *emr.Client
	Emrcontainersconn                   *emrcontainers.Client
	Eventbridgeconn                     *eventbridge.Client
	Firehoseconn                        *firehose.Client
	Fmsconn                             *fms.Client
	Frauddetectorconn                   *frauddetector.Client
	Fsxconn                             *fsx.Client
	Gameliftconn                        *gamelift.Client
	Glacierconn                         *glacier.Client
	Globalacceleratorconn               *globalaccelerator.Client
	Glueconn                            *glue.Client
	Greengrassconn                      *greengrass.Client
	Greengrassv2conn                    *greengrassv2.Client
	Groundstationconn                   *groundstation.Client
	Guarddutyconn                       *guardduty.Client
	Healthconn                          *health.Client
	Healthlakeconn                      *healthlake.Client
	Honeycodeconn                       *honeycode.Client
	Iamconn                             *iam.Client
	Identitystoreconn                   *identitystore.Client
	Imagebuilderconn                    *imagebuilder.Client
	Inspectorconn                       *inspector.Client
	Iotconn                             *iot.Client
	Iot1clickdevicesserviceconn         *iot1clickdevicesservice.Client
	Iot1clickprojectsconn               *iot1clickprojects.Client
	Iotanalyticsconn                    *iotanalytics.Client
	Iotdataplaneconn                    *iotdataplane.Client
	Iotdeviceadvisorconn                *iotdeviceadvisor.Client
	Ioteventsconn                       *iotevents.Client
	Ioteventsdataconn                   *ioteventsdata.Client
	Iotfleethubconn                     *iotfleethub.Client
	Iotjobsdataplaneconn                *iotjobsdataplane.Client
	Iotsecuretunnelingconn              *iotsecuretunneling.Client
	Iotsitewiseconn                     *iotsitewise.Client
	Iotthingsgraphconn                  *iotthingsgraph.Client
	Iotwirelessconn                     *iotwireless.Client
	Ivsconn                             *ivs.Client
	Kafkaconn                           *kafka.Client
	Kendraconn                          *kendra.Client
	Kinesisconn                         *kinesis.Client
	Kinesisanalyticsconn                *kinesisanalytics.Client
	Kinesisanalyticsv2conn              *kinesisanalyticsv2.Client
	Kinesisvideoconn                    *kinesisvideo.Client
	Kinesisvideoarchivedmediaconn       *kinesisvideoarchivedmedia.Client
	Kinesisvideomediaconn               *kinesisvideomedia.Client
	Kmsconn                             *kms.Client
	Lakeformationconn                   *lakeformation.Client
	Lambdaconn                          *lambda.Client
	Lexmodelbuildingserviceconn         *lexmodelbuildingservice.Client
	Lexruntimeserviceconn               *lexruntimeservice.Client
	Licensemanagerconn                  *licensemanager.Client
	Lightsailconn                       *lightsail.Client
	Machinelearningconn                 *machinelearning.Client
	Macieconn                           *macie.Client
	Macie2conn                          *macie2.Client
	Managedblockchainconn               *managedblockchain.Client
	Marketplacecatalogconn              *marketplacecatalog.Client
	Marketplacecommerceanalyticsconn    *marketplacecommerceanalytics.Client
	Marketplaceentitlementserviceconn   *marketplaceentitlementservice.Client
	Marketplacemeteringconn             *marketplacemetering.Client
	Mediaconnectconn                    *mediaconnect.Client
	Mediaconvertconn                    *mediaconvert.Client
	Medialiveconn                       *medialive.Client
	Mediapackageconn                    *mediapackage.Client
	Mediapackagevodconn                 *mediapackagevod.Client
	Mediastoreconn                      *mediastore.Client
	Mediastoredataconn                  *mediastoredata.Client
	Mediatailorconn                     *mediatailor.Client
	Migrationhubconn                    *migrationhub.Client
	Migrationhubconfigconn              *migrationhubconfig.Client
	Mobileconn                          *mobile.Client
	Mqconn                              *mq.Client
	Mturkconn                           *mturk.Client
	Neptuneconn                         *neptune.Client
	Networkfirewallconn                 *networkfirewall.Client
	Networkmanagerconn                  *networkmanager.Client
	Opsworksconn                        *opsworks.Client
	Opsworkscmconn                      *opsworkscm.Client
	Organizationsconn                   *organizations.Client
	Outpostsconn                        *outposts.Client
	Personalizeconn                     *personalize.Client
	Personalizeeventsconn               *personalizeevents.Client
	Personalizeruntimeconn              *personalizeruntime.Client
	Piconn                              *pi.Client
	Pinpointconn                        *pinpoint.Client
	Pinpointemailconn                   *pinpointemail.Client
	Pinpointsmsvoiceconn                *pinpointsmsvoice.Client
	Pollyconn                           *polly.Client
	Pricingconn                         *pricing.Client
	Qldbconn                            *qldb.Client
	Qldbsessionconn                     *qldbsession.Client
	Quicksightconn                      *quicksight.Client
	Ramconn                             *ram.Client
	Rdsconn                             *rds.Client
	Redshiftconn                        *redshift.Client
	Rekognitionconn                     *rekognition.Client
	Resourcegroupsconn                  *resourcegroups.Client
	Resourcegroupstaggingapiconn        *resourcegroupstaggingapi.Client
	Robomakerconn                       *robomaker.Client
	Route53conn                         *route53.Client
	Route53domainsconn                  *route53domains.Client
	Route53resolverconn                 *route53resolver.Client
	S3conn                              *s3.Client
	S3controlconn                       *s3control.Client
	S3outpostsconn                      *s3outposts.Client
	Sagemakerconn                       *sagemaker.Client
	Sagemakerfeaturestoreruntimeconn    *sagemakerfeaturestoreruntime.Client
	Sagemakerruntimeconn                *sagemakerruntime.Client
	Savingsplansconn                    *savingsplans.Client
	Schemasconn                         *schemas.Client
	Secretsmanagerconn                  *secretsmanager.Client
	Securityhubconn                     *securityhub.Client
	Serverlessapplicationrepositoryconn *serverlessapplicationrepository.Client
	Servicecatalogconn                  *servicecatalog.Client
	Servicediscoveryconn                *servicediscovery.Client
	Servicequotasconn                   *servicequotas.Client
	Sesconn                             *ses.Client
	Sesv2conn                           *sesv2.Client
	Sfnconn                             *sfn.Client
	Shieldconn                          *shield.Client
	Signerconn                          *signer.Client
	Smsconn                             *sms.Client
	Snowballconn                        *snowball.Client
	Snsconn                             *sns.Client
	Sqsconn                             *sqs.Client
	Ssmconn                             *ssm.Client
	Ssoconn                             *sso.Client
	Ssoadminconn                        *ssoadmin.Client
	Ssooidcconn                         *ssooidc.Client
	Storagegatewayconn                  *storagegateway.Client
	Stsconn                             *sts.Client
	Supportconn                         *support.Client
	Swfconn                             *swf.Client
	Syntheticsconn                      *synthetics.Client
	Textractconn                        *textract.Client
	Timestreamqueryconn                 *timestreamquery.Client
	Timestreamwriteconn                 *timestreamwrite.Client
	Transferconn                        *transfer.Client
	Translateconn                       *translate.Client
	Wafconn                             *waf.Client
	Wafregionalconn                     *wafregional.Client
	Wafv2conn                           *wafv2.Client
	Wellarchitectedconn                 *wellarchitected.Client
	Workdocsconn                        *workdocs.Client
	Worklinkconn                        *worklink.Client
	Workmailconn                        *workmail.Client
	Workmailmessageflowconn             *workmailmessageflow.Client
	Workspacesconn                      *workspaces.Client
	Xrayconn                            *xray.Client
}

func NewClient(ctx context.Context, configs ...func(*config.LoadOptions) error) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		Accessanalyzerconn:                  accessanalyzer.NewFromConfig(cfg),
		Acmconn:                             acm.NewFromConfig(cfg),
		Acmpcaconn:                          acmpca.NewFromConfig(cfg),
		Alexaforbusinessconn:                alexaforbusiness.NewFromConfig(cfg),
		Amplifyconn:                         amplify.NewFromConfig(cfg),
		Apigatewayconn:                      apigateway.NewFromConfig(cfg),
		Apigatewaymanagementapiconn:         apigatewaymanagementapi.NewFromConfig(cfg),
		Apigatewayv2conn:                    apigatewayv2.NewFromConfig(cfg),
		Appconfigconn:                       appconfig.NewFromConfig(cfg),
		Appflowconn:                         appflow.NewFromConfig(cfg),
		Applicationautoscalingconn:          applicationautoscaling.NewFromConfig(cfg),
		Applicationdiscoveryserviceconn:     applicationdiscoveryservice.NewFromConfig(cfg),
		Applicationinsightsconn:             applicationinsights.NewFromConfig(cfg),
		Appmeshconn:                         appmesh.NewFromConfig(cfg),
		Appstreamconn:                       appstream.NewFromConfig(cfg),
		Appsyncconn:                         appsync.NewFromConfig(cfg),
		Athenaconn:                          athena.NewFromConfig(cfg),
		Auditmanagerconn:                    auditmanager.NewFromConfig(cfg),
		Autoscalingconn:                     autoscaling.NewFromConfig(cfg),
		Autoscalingplansconn:                autoscalingplans.NewFromConfig(cfg),
		Backupconn:                          backup.NewFromConfig(cfg),
		Batchconn:                           batch.NewFromConfig(cfg),
		Braketconn:                          braket.NewFromConfig(cfg),
		Budgetsconn:                         budgets.NewFromConfig(cfg),
		Chimeconn:                           chime.NewFromConfig(cfg),
		Cloud9conn:                          cloud9.NewFromConfig(cfg),
		Clouddirectoryconn:                  clouddirectory.NewFromConfig(cfg),
		Cloudformationconn:                  cloudformation.NewFromConfig(cfg),
		Cloudfrontconn:                      cloudfront.NewFromConfig(cfg),
		Cloudhsmconn:                        cloudhsm.NewFromConfig(cfg),
		Cloudhsmv2conn:                      cloudhsmv2.NewFromConfig(cfg),
		Cloudsearchconn:                     cloudsearch.NewFromConfig(cfg),
		Cloudsearchdomainconn:               cloudsearchdomain.NewFromConfig(cfg),
		Cloudtrailconn:                      cloudtrail.NewFromConfig(cfg),
		Cloudwatchconn:                      cloudwatch.NewFromConfig(cfg),
		Cloudwatcheventsconn:                cloudwatchevents.NewFromConfig(cfg),
		Cloudwatchlogsconn:                  cloudwatchlogs.NewFromConfig(cfg),
		Codeartifactconn:                    codeartifact.NewFromConfig(cfg),
		Codebuildconn:                       codebuild.NewFromConfig(cfg),
		Codecommitconn:                      codecommit.NewFromConfig(cfg),
		Codedeployconn:                      codedeploy.NewFromConfig(cfg),
		Codeguruprofilerconn:                codeguruprofiler.NewFromConfig(cfg),
		Codegurureviewerconn:                codegurureviewer.NewFromConfig(cfg),
		Codepipelineconn:                    codepipeline.NewFromConfig(cfg),
		Codestarconn:                        codestar.NewFromConfig(cfg),
		Codestarconnectionsconn:             codestarconnections.NewFromConfig(cfg),
		Codestarnotificationsconn:           codestarnotifications.NewFromConfig(cfg),
		Cognitoidentityconn:                 cognitoidentity.NewFromConfig(cfg),
		Cognitoidentityproviderconn:         cognitoidentityprovider.NewFromConfig(cfg),
		Cognitosyncconn:                     cognitosync.NewFromConfig(cfg),
		Comprehendconn:                      comprehend.NewFromConfig(cfg),
		Comprehendmedicalconn:               comprehendmedical.NewFromConfig(cfg),
		Computeoptimizerconn:                computeoptimizer.NewFromConfig(cfg),
		Configserviceconn:                   configservice.NewFromConfig(cfg),
		Connectconn:                         connect.NewFromConfig(cfg),
		Connectcontactlensconn:              connectcontactlens.NewFromConfig(cfg),
		Connectparticipantconn:              connectparticipant.NewFromConfig(cfg),
		Costandusagereportserviceconn:       costandusagereportservice.NewFromConfig(cfg),
		Costexplorerconn:                    costexplorer.NewFromConfig(cfg),
		Customerprofilesconn:                customerprofiles.NewFromConfig(cfg),
		Databasemigrationserviceconn:        databasemigrationservice.NewFromConfig(cfg),
		Dataexchangeconn:                    dataexchange.NewFromConfig(cfg),
		Datapipelineconn:                    datapipeline.NewFromConfig(cfg),
		Datasyncconn:                        datasync.NewFromConfig(cfg),
		Daxconn:                             dax.NewFromConfig(cfg),
		Detectiveconn:                       detective.NewFromConfig(cfg),
		Devicefarmconn:                      devicefarm.NewFromConfig(cfg),
		Devopsguruconn:                      devopsguru.NewFromConfig(cfg),
		Directconnectconn:                   directconnect.NewFromConfig(cfg),
		Directoryserviceconn:                directoryservice.NewFromConfig(cfg),
		Dlmconn:                             dlm.NewFromConfig(cfg),
		Docdbconn:                           docdb.NewFromConfig(cfg),
		Dynamodbconn:                        dynamodb.NewFromConfig(cfg),
		Dynamodbstreamsconn:                 dynamodbstreams.NewFromConfig(cfg),
		Ebsconn:                             ebs.NewFromConfig(cfg),
		Ec2conn:                             ec2.NewFromConfig(cfg),
		Ec2instanceconnectconn:              ec2instanceconnect.NewFromConfig(cfg),
		Ecrconn:                             ecr.NewFromConfig(cfg),
		Ecrpublicconn:                       ecrpublic.NewFromConfig(cfg),
		Ecsconn:                             ecs.NewFromConfig(cfg),
		Efsconn:                             efs.NewFromConfig(cfg),
		Eksconn:                             eks.NewFromConfig(cfg),
		Elasticacheconn:                     elasticache.NewFromConfig(cfg),
		Elasticbeanstalkconn:                elasticbeanstalk.NewFromConfig(cfg),
		Elasticinferenceconn:                elasticinference.NewFromConfig(cfg),
		Elasticloadbalancingconn:            elasticloadbalancing.NewFromConfig(cfg),
		Elasticloadbalancingv2conn:          elasticloadbalancingv2.NewFromConfig(cfg),
		Elasticsearchserviceconn:            elasticsearchservice.NewFromConfig(cfg),
		Elastictranscoderconn:               elastictranscoder.NewFromConfig(cfg),
		Emrconn:                             emr.NewFromConfig(cfg),
		Emrcontainersconn:                   emrcontainers.NewFromConfig(cfg),
		Eventbridgeconn:                     eventbridge.NewFromConfig(cfg),
		Firehoseconn:                        firehose.NewFromConfig(cfg),
		Fmsconn:                             fms.NewFromConfig(cfg),
		Frauddetectorconn:                   frauddetector.NewFromConfig(cfg),
		Fsxconn:                             fsx.NewFromConfig(cfg),
		Gameliftconn:                        gamelift.NewFromConfig(cfg),
		Glacierconn:                         glacier.NewFromConfig(cfg),
		Globalacceleratorconn:               globalaccelerator.NewFromConfig(cfg),
		Glueconn:                            glue.NewFromConfig(cfg),
		Greengrassconn:                      greengrass.NewFromConfig(cfg),
		Greengrassv2conn:                    greengrassv2.NewFromConfig(cfg),
		Groundstationconn:                   groundstation.NewFromConfig(cfg),
		Guarddutyconn:                       guardduty.NewFromConfig(cfg),
		Healthconn:                          health.NewFromConfig(cfg),
		Healthlakeconn:                      healthlake.NewFromConfig(cfg),
		Honeycodeconn:                       honeycode.NewFromConfig(cfg),
		Iamconn:                             iam.NewFromConfig(cfg),
		Identitystoreconn:                   identitystore.NewFromConfig(cfg),
		Imagebuilderconn:                    imagebuilder.NewFromConfig(cfg),
		Inspectorconn:                       inspector.NewFromConfig(cfg),
		Iotconn:                             iot.NewFromConfig(cfg),
		Iot1clickdevicesserviceconn:         iot1clickdevicesservice.NewFromConfig(cfg),
		Iot1clickprojectsconn:               iot1clickprojects.NewFromConfig(cfg),
		Iotanalyticsconn:                    iotanalytics.NewFromConfig(cfg),
		Iotdataplaneconn:                    iotdataplane.NewFromConfig(cfg),
		Iotdeviceadvisorconn:                iotdeviceadvisor.NewFromConfig(cfg),
		Ioteventsconn:                       iotevents.NewFromConfig(cfg),
		Ioteventsdataconn:                   ioteventsdata.NewFromConfig(cfg),
		Iotfleethubconn:                     iotfleethub.NewFromConfig(cfg),
		Iotjobsdataplaneconn:                iotjobsdataplane.NewFromConfig(cfg),
		Iotsecuretunnelingconn:              iotsecuretunneling.NewFromConfig(cfg),
		Iotsitewiseconn:                     iotsitewise.NewFromConfig(cfg),
		Iotthingsgraphconn:                  iotthingsgraph.NewFromConfig(cfg),
		Iotwirelessconn:                     iotwireless.NewFromConfig(cfg),
		Ivsconn:                             ivs.NewFromConfig(cfg),
		Kafkaconn:                           kafka.NewFromConfig(cfg),
		Kendraconn:                          kendra.NewFromConfig(cfg),
		Kinesisconn:                         kinesis.NewFromConfig(cfg),
		Kinesisanalyticsconn:                kinesisanalytics.NewFromConfig(cfg),
		Kinesisanalyticsv2conn:              kinesisanalyticsv2.NewFromConfig(cfg),
		Kinesisvideoconn:                    kinesisvideo.NewFromConfig(cfg),
		Kinesisvideoarchivedmediaconn:       kinesisvideoarchivedmedia.NewFromConfig(cfg),
		Kinesisvideomediaconn:               kinesisvideomedia.NewFromConfig(cfg),
		Kmsconn:                             kms.NewFromConfig(cfg),
		Lakeformationconn:                   lakeformation.NewFromConfig(cfg),
		Lambdaconn:                          lambda.NewFromConfig(cfg),
		Lexmodelbuildingserviceconn:         lexmodelbuildingservice.NewFromConfig(cfg),
		Lexruntimeserviceconn:               lexruntimeservice.NewFromConfig(cfg),
		Licensemanagerconn:                  licensemanager.NewFromConfig(cfg),
		Lightsailconn:                       lightsail.NewFromConfig(cfg),
		Machinelearningconn:                 machinelearning.NewFromConfig(cfg),
		Macieconn:                           macie.NewFromConfig(cfg),
		Macie2conn:                          macie2.NewFromConfig(cfg),
		Managedblockchainconn:               managedblockchain.NewFromConfig(cfg),
		Marketplacecatalogconn:              marketplacecatalog.NewFromConfig(cfg),
		Marketplacecommerceanalyticsconn:    marketplacecommerceanalytics.NewFromConfig(cfg),
		Marketplaceentitlementserviceconn:   marketplaceentitlementservice.NewFromConfig(cfg),
		Marketplacemeteringconn:             marketplacemetering.NewFromConfig(cfg),
		Mediaconnectconn:                    mediaconnect.NewFromConfig(cfg),
		Mediaconvertconn:                    mediaconvert.NewFromConfig(cfg),
		Medialiveconn:                       medialive.NewFromConfig(cfg),
		Mediapackageconn:                    mediapackage.NewFromConfig(cfg),
		Mediapackagevodconn:                 mediapackagevod.NewFromConfig(cfg),
		Mediastoreconn:                      mediastore.NewFromConfig(cfg),
		Mediastoredataconn:                  mediastoredata.NewFromConfig(cfg),
		Mediatailorconn:                     mediatailor.NewFromConfig(cfg),
		Migrationhubconn:                    migrationhub.NewFromConfig(cfg),
		Migrationhubconfigconn:              migrationhubconfig.NewFromConfig(cfg),
		Mobileconn:                          mobile.NewFromConfig(cfg),
		Mqconn:                              mq.NewFromConfig(cfg),
		Mturkconn:                           mturk.NewFromConfig(cfg),
		Neptuneconn:                         neptune.NewFromConfig(cfg),
		Networkfirewallconn:                 networkfirewall.NewFromConfig(cfg),
		Networkmanagerconn:                  networkmanager.NewFromConfig(cfg),
		Opsworksconn:                        opsworks.NewFromConfig(cfg),
		Opsworkscmconn:                      opsworkscm.NewFromConfig(cfg),
		Organizationsconn:                   organizations.NewFromConfig(cfg),
		Outpostsconn:                        outposts.NewFromConfig(cfg),
		Personalizeconn:                     personalize.NewFromConfig(cfg),
		Personalizeeventsconn:               personalizeevents.NewFromConfig(cfg),
		Personalizeruntimeconn:              personalizeruntime.NewFromConfig(cfg),
		Piconn:                              pi.NewFromConfig(cfg),
		Pinpointconn:                        pinpoint.NewFromConfig(cfg),
		Pinpointemailconn:                   pinpointemail.NewFromConfig(cfg),
		Pinpointsmsvoiceconn:                pinpointsmsvoice.NewFromConfig(cfg),
		Pollyconn:                           polly.NewFromConfig(cfg),
		Pricingconn:                         pricing.NewFromConfig(cfg),
		Qldbconn:                            qldb.NewFromConfig(cfg),
		Qldbsessionconn:                     qldbsession.NewFromConfig(cfg),
		Quicksightconn:                      quicksight.NewFromConfig(cfg),
		Ramconn:                             ram.NewFromConfig(cfg),
		Rdsconn:                             rds.NewFromConfig(cfg),
		Redshiftconn:                        redshift.NewFromConfig(cfg),
		Rekognitionconn:                     rekognition.NewFromConfig(cfg),
		Resourcegroupsconn:                  resourcegroups.NewFromConfig(cfg),
		Resourcegroupstaggingapiconn:        resourcegroupstaggingapi.NewFromConfig(cfg),
		Robomakerconn:                       robomaker.NewFromConfig(cfg),
		Route53conn:                         route53.NewFromConfig(cfg),
		Route53domainsconn:                  route53domains.NewFromConfig(cfg),
		Route53resolverconn:                 route53resolver.NewFromConfig(cfg),
		S3conn:                              s3.NewFromConfig(cfg),
		S3controlconn:                       s3control.NewFromConfig(cfg),
		S3outpostsconn:                      s3outposts.NewFromConfig(cfg),
		Sagemakerconn:                       sagemaker.NewFromConfig(cfg),
		Sagemakerfeaturestoreruntimeconn:    sagemakerfeaturestoreruntime.NewFromConfig(cfg),
		Sagemakerruntimeconn:                sagemakerruntime.NewFromConfig(cfg),
		Savingsplansconn:                    savingsplans.NewFromConfig(cfg),
		Schemasconn:                         schemas.NewFromConfig(cfg),
		Secretsmanagerconn:                  secretsmanager.NewFromConfig(cfg),
		Securityhubconn:                     securityhub.NewFromConfig(cfg),
		Serverlessapplicationrepositoryconn: serverlessapplicationrepository.NewFromConfig(cfg),
		Servicecatalogconn:                  servicecatalog.NewFromConfig(cfg),
		Servicediscoveryconn:                servicediscovery.NewFromConfig(cfg),
		Servicequotasconn:                   servicequotas.NewFromConfig(cfg),
		Sesconn:                             ses.NewFromConfig(cfg),
		Sesv2conn:                           sesv2.NewFromConfig(cfg),
		Sfnconn:                             sfn.NewFromConfig(cfg),
		Shieldconn:                          shield.NewFromConfig(cfg),
		Signerconn:                          signer.NewFromConfig(cfg),
		Smsconn:                             sms.NewFromConfig(cfg),
		Snowballconn:                        snowball.NewFromConfig(cfg),
		Snsconn:                             sns.NewFromConfig(cfg),
		Sqsconn:                             sqs.NewFromConfig(cfg),
		Ssmconn:                             ssm.NewFromConfig(cfg),
		Ssoconn:                             sso.NewFromConfig(cfg),
		Ssoadminconn:                        ssoadmin.NewFromConfig(cfg),
		Ssooidcconn:                         ssooidc.NewFromConfig(cfg),
		Storagegatewayconn:                  storagegateway.NewFromConfig(cfg),
		Stsconn:                             sts.NewFromConfig(cfg),
		Supportconn:                         support.NewFromConfig(cfg),
		Swfconn:                             swf.NewFromConfig(cfg),
		Syntheticsconn:                      synthetics.NewFromConfig(cfg),
		Textractconn:                        textract.NewFromConfig(cfg),
		Timestreamqueryconn:                 timestreamquery.NewFromConfig(cfg),
		Timestreamwriteconn:                 timestreamwrite.NewFromConfig(cfg),
		Transferconn:                        transfer.NewFromConfig(cfg),
		Translateconn:                       translate.NewFromConfig(cfg),
		Wafconn:                             waf.NewFromConfig(cfg),
		Wafregionalconn:                     wafregional.NewFromConfig(cfg),
		Wafv2conn:                           wafv2.NewFromConfig(cfg),
		Wellarchitectedconn:                 wellarchitected.NewFromConfig(cfg),
		Workdocsconn:                        workdocs.NewFromConfig(cfg),
		Worklinkconn:                        worklink.NewFromConfig(cfg),
		Workmailconn:                        workmail.NewFromConfig(cfg),
		Workmailmessageflowconn:             workmailmessageflow.NewFromConfig(cfg),
		Workspacesconn:                      workspaces.NewFromConfig(cfg),
		Xrayconn:                            xray.NewFromConfig(cfg),
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
		return fmt.Errorf("failed to get caller identity: %s", err)
	}

	client.AccountID = *resp.Account

	return nil
}
