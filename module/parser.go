package module

import (
	"fmt"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func ParseModule(path string) *tfconfig.Module {
	module, _ := tfconfig.LoadModule(path)
	//ctx.BindJSON(&newModule)
	for key, val := range module.Variables {
		fmt.Println(key, "required: ", val.Required)
	}

	return module
}
