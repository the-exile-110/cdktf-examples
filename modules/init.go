package modules

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type Module struct {
	scope constructs.Construct
}

func NewModule(scope constructs.Construct) *Module {
	return &Module{scope: scope}
}
