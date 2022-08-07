package stacks

import (
	"cdktf-examples/modules"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s Stack) Ec2Stack() cdktf.TerraformStack {
	modules.NewModule(s.stack).Ec2Instance("testEc2", &ec2.InstanceConfig{
		Ami:          jsii.String("ami-2757f631"),
		InstanceType: jsii.String("t2.micro"),
		Tags: &map[string]*string{
			"environment": jsii.String("development"),
		},
	})

	return s.stack
}
