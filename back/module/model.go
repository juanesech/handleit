package module

import "github.com/hashicorp/terraform-config-inspect/tfconfig"

type ModuleResume struct {
	Name      string
	Providers map[string]*tfconfig.ProviderRequirement
}

type Module struct {
	ID        string
	Name      string
	Variables []Variable
	Outputs   []Output
	Providers map[string]*tfconfig.ProviderRequirement
}

type Variable struct {
	Name        string
	Type        string
	Description string
	Default     string
	Required    bool
}

type Output struct {
	Name        string
	Description string
}

type ImportRequest struct {
	Name string `json:"name"`
}
