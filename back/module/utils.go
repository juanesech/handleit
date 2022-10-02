package module

import (
	"fmt"
	"os"
	"path/filepath"

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
