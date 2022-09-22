package module

import "github.com/hashicorp/terraform-config-inspect/tfconfig"

type ModuleResume struct {
	Name      string
	Providers map[string]*tfconfig.ProviderRequirement
}

type Module struct {
	ID        string
	Name      string
	Variables map[string]*tfconfig.Variable
	Outputs   map[string]*tfconfig.Output
	Providers map[string]*tfconfig.ProviderRequirement
}

type ImportRequest struct {
	Name string `json:"name"`
}
