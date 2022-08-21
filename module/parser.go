package module

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func ParseModule(path string) *tfconfig.Module {
	module, _ := tfconfig.LoadModule(path)

	return module
}
