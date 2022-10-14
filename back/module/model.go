package module

type ModuleResume struct {
	Name      string
	Providers []Provider
}

type Module struct {
	ID        string
	Name      string
	Variables []Variable
	Outputs   []Output
	Providers []Provider
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

type Provider struct {
	Source             string
	VersionConstraints []string
}

type ImportRequest struct {
	Name string `json:"name"`
}
