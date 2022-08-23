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
	db "github.com/juanesech/handleit/database"
)

type ModuleResume struct {
	Name      string
	Providers map[string]*tfconfig.ProviderRequirement
}

type Module struct {
	Name      string
	Variables map[string]*tfconfig.Variable
	Outputs   map[string]*tfconfig.Output
	Providers map[string]*tfconfig.ProviderRequirement
}

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

func List(ctx *gin.Context) {
	var moduleList []ModuleResume

	for _, module := range getModulesFromFS() {
		moduleResume := &ModuleResume{
			Name:      module.Name,
			Providers: module.Providers,
		}
		moduleList = append(moduleList, *moduleResume)
	}
	ctx.JSON(http.StatusOK, moduleList)
}

func Get(ctx *gin.Context) {
	// moduleName := ctx.Param("name")
	// module := ParseModule(fmt.Sprintf("%s/%s", config.Get().ModuleSource, moduleName))

	ctx.JSON(http.StatusOK, "WIP")
}

func Import(ctx *gin.Context) {
	var modulesFromFS []*Module = getModulesFromFS()
	session, err := db.Client.OpenSession("handleit")
	if err != nil {
		log.Error("Can't open database session ", err)
	}

	for _, module := range modulesFromFS {
		adv := session.Advanced()
		log.Info("Document ID: ", adv.GetDocumentID(module))
		err = session.Store(module)
		if err != nil {
			log.Error("Failed to save ", err)
		}
	}
	err = session.SaveChanges()
	if err != nil {
		log.Error("Failed to save ", err)
	}
}
