package module

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/juanesech/handleit/config"
	log "github.com/sirupsen/logrus"
)

func GetModuleName(path string) string {
	return filepath.Base(path)
}

func getModulesFromFS() []*Module {
	var moduleList []*Module
	files, err := ioutil.ReadDir(config.Get().ModuleSource)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		modulePath := fmt.Sprintf("%s/%s", config.Get().ModuleSource, f.Name())
		moduleList = append(moduleList, ParseModule(modulePath))
	}

	return moduleList
}
