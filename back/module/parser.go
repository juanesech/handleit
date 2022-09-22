package module

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func ParseModule(path string) *Module {
	var parsedModule *Module
	module, _ := tfconfig.LoadModule(path)
	parsedModule = &Module{
		Name:      GetModuleName(path),
		Variables: module.Variables,
		Outputs:   module.Outputs,
		Providers: module.RequiredProviders,
	}
	return parsedModule
}
