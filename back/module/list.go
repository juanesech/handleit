package module

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/juanesech/topo/constants"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
)

func List(ctx *gin.Context) {
	var moduleList []ModuleResume
	var modulesFromDB []*Module

	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	query := session.QueryCollectionForType(reflect.TypeOf(&Module{}))
	utils.CheckError(query.GetResults(&modulesFromDB))

	for _, module := range modulesFromDB {
		moduleResume := &ModuleResume{
			Name:      module.Name,
			Providers: module.Providers,
		}
		moduleList = append(moduleList, *moduleResume)
	}
	ctx.JSON(http.StatusOK, moduleList)
}
