package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (m *Module) Ec2Instance(id string, instanceConfig *ec2.InstanceConfig) ec2.Instance {
	instance := ec2.NewInstance(m.scope, jsii.String(id), instanceConfig)
	cdktf.NewTerraformOutput(m.scope, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return instance
}

func (m *Module) Eip(id string, eipConfig *ec2.EipConfig) ec2.Eip {
	eip := ec2.NewEip(m.scope, jsii.String(id), eipConfig)
	return eip
}
