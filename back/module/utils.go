package module

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/juanesech/topo/utils"
)

func GetModuleName(path string) string {
	return filepath.Base(path)
}

func getModulesFromFS(path string) []*Module {
	var moduleList []*Module
	files, err := os.ReadDir(path)
	utils.CheckError(err)

	for _, f := range files {
		modulePath := fmt.Sprintf("%s/%s", path, f.Name())
		moduleList = append(moduleList, ParseModule(modulePath))
	}

	return moduleList
}

func varToArray(vars map[string]*tfconfig.Variable) []Variable {
	array := []Variable{}
	var va = Variable{}
	for k, v := range vars {
		va.Name = k
		va.Description = v.Description
		va.Type = v.Type
		va.Default = fmt.Sprintf(`%v`, v.Default)
		va.Required = v.Required

		array = append(array, va)
	}
	return array
}

func outToArray(vars map[string]*tfconfig.Output) []Output {
	array := []Output{}
	var ou = Output{}
	for k, v := range vars {
		ou.Name = k
		ou.Description = v.Description

		array = append(array, ou)
	}
	return array
}
