module github.com/jckuester/awstools-lib

go 1.15

require (
	github.com/apex/log v1.9.0
	github.com/aws/aws-sdk-go v1.35.28
	github.com/aws/aws-sdk-go-v2/config v1.1.1
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/acm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.1.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.1.1
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.1.1
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.1.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.1.1
	github.com/aws/aws-sdk-go-v2/service/athena v1.1.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.1.1
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.1.1
	github.com/aws/aws-sdk-go-v2/service/backup v1.1.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.1.1
	github.com/aws/aws-sdk-go-v2/service/budgets v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.1.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.1.1
	github.com/aws/aws-sdk-go-v2/service/datasync v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.1.1
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.1.1
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dlm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/docdb v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.1.2
	github.com/aws/aws-sdk-go-v2/service/ecs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.1.1
	github.com/aws/aws-sdk-go-v2/service/emr v1.1.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.1.1
	github.com/aws/aws-sdk-go-v2/service/fms v1.1.1
	github.com/aws/aws-sdk-go-v2/service/fsx v1.1.1
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.1.1
	github.com/aws/aws-sdk-go-v2/service/glacier v1.1.1
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.1.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.1.1
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.1.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.1.1
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.1.1
	github.com/aws/aws-sdk-go-v2/service/inspector v1.1.1
	github.com/aws/aws-sdk-go-v2/service/iot v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kafka v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.1.1
	github.com/aws/aws-sdk-go-v2/service/macie v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mq v1.1.1
	github.com/aws/aws-sdk-go-v2/service/neptune v1.1.1
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.1.2
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.1.1
	github.com/aws/aws-sdk-go-v2/service/organizations v1.1.1
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.1.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.1
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ram v1.1.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.1.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.1.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.1.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.1.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.2.0
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.1.2
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.1.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.1.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.1.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.1.1
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.1.1
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.1.1
	github.com/aws/aws-sdk-go-v2/service/shield v1.1.1
	github.com/aws/aws-sdk-go-v2/service/signer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ssm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.1.2
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.1.1
	github.com/aws/aws-sdk-go-v2/service/swf v1.1.1
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.1.1
	github.com/aws/aws-sdk-go-v2/service/transfer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.1
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.1.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/worklink v1.1.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.1.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.1.1
	github.com/fatih/color v1.10.0
	github.com/jckuester/awsls v0.8.3
	github.com/jckuester/terradozer v0.1.4-0.20210524132110-395dc12d3bf5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/zclconf/go-cty v1.7.1
)
