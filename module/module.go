package module

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/juanesech/handleit/config"
)

//Just for testing proposes
var Database []tfconfig.Module

func List(ctx *gin.Context) {
	files, err := ioutil.ReadDir(config.Conf.ModuleSource)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		modulePath := fmt.Sprintf("%s/%s", config.Conf.ModuleSource, f.Name())
		Database = append(Database, *ParseModule(modulePath))
	}
	ctx.JSON(http.StatusOK, Database)
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
	ctx.JSON(http.StatusOK, module)
}
