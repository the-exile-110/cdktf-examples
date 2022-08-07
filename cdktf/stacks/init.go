package stacks

import (
	"cdk.tf/go/stack/config"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Stack struct {
	cdktf.TerraformStack
}

func NewStack(scope constructs.Construct, id string) *Stack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String(config.Region),
	})
	return &Stack{stack}
}
