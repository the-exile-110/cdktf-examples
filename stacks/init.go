package stacks

import (
	"cdktf-examples/config"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Stack struct {
	stack cdktf.TerraformStack
}

func NewStack(scope constructs.Construct, id string, stateKey string) *Stack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String(config.Region),
	})
	cdktf.NewS3Backend(stack, &cdktf.S3BackendProps{
		Region: jsii.String(config.Region),
		Bucket: jsii.String(config.BackendBucket),
		Key:    jsii.String(stateKey),
	})
	return &Stack{stack: stack}
}
