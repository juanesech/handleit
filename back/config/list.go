package config

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/juanesech/topo/constants"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
)

func List(ctx *gin.Context) {
//	var sourceList []ModuleSource
	var sourcesFromDB []*ModuleSource

	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	query := session.QueryCollectionForType(reflect.TypeOf(&ModuleSource{}))
	utils.CheckError(query.GetResults(&sourcesFromDB))

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, sourcesFromDB)
}