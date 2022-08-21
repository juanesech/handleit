package module

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/juanesech/handleit/config"
)

//Just for testing proposes
var Database []tfconfig.Module

type ModuleResume struct {
	Name      string
	Providers map[string]*tfconfig.ProviderRequirement
}

func GetModuleName(path string) string {
	return filepath.Base(path)
}

func GetModulesFromFS() []tfconfig.Module {
	var moduleList []tfconfig.Module
	files, err := ioutil.ReadDir(config.Conf.ModuleSource)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		modulePath := fmt.Sprintf("%s/%s", config.Conf.ModuleSource, f.Name())
		moduleList = append(moduleList, *ParseModule(modulePath))
	}

	return moduleList
}

func List(ctx *gin.Context) {
	var moduleList []ModuleResume

	for _, module := range GetModulesFromFS() {
		moduleResume := &ModuleResume{
			Name:      GetModuleName(module.Path),
			Providers: module.RequiredProviders,
		}
		moduleList = append(moduleList, *moduleResume)
	}
	ctx.JSON(http.StatusOK, moduleList)
}

func New(ctx *gin.Context) {
	var newModule tfconfig.Module
	ctx.BindJSON(&newModule)
	Database = append(Database, newModule)
	message := fmt.Sprintf("Module %s posted", newModule.Path)
	ctx.String(http.StatusOK, message)
}

func Get(ctx *gin.Context) {
	moduleName := ctx.Param("name")
	module := ParseModule(fmt.Sprintf("%s/%s", config.Conf.ModuleSource, moduleName))

	if module.Diagnostics != nil {
		ctx.JSON(http.StatusNotFound, module.Diagnostics)
		log.Error(module.Diagnostics.Error())
	} else {
		ctx.JSON(http.StatusOK, module)
	}
}
