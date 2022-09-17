package module

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func GetModuleName(path string) string {
	return filepath.Base(path)
}

func getModulesFromFS(path string) []*Module {
	var moduleList []*Module
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		modulePath := fmt.Sprintf("%s/%s", path, f.Name())
		moduleList = append(moduleList, ParseModule(modulePath))
	}

	return moduleList
}
