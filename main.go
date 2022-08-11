package main

import (
	"cdktf-examples/stacks"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	stacks.NewStack(app, "ec2Stack", "ec2/ec2.tfstate").Ec2Stack()

	app.Synth()
}
