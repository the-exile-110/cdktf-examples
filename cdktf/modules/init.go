package modules

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type Module struct {
	stack *cdktf.TerraformStack
}

func NewModule(stack *cdktf.TerraformStack) *Module {
	return &Module{stack: stack}
}
