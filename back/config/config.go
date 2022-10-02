package config

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/juanesech/topo/constants"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Sources []ModuleSource
}

type ModuleSource struct {
	ID      string
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Group   string `json:"group"`
	Auth    string `json:"auth"`
}

func Set(ctx *gin.Context) {
	var sourceFromReq *ModuleSource
	var sourcesFromDB []*ModuleSource
	var source *ModuleSource

	ctx.BindJSON(&sourceFromReq)

	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	query := session.QueryCollectionForType(reflect.TypeOf(&ModuleSource{})).WhereEquals("name", sourceFromReq.Name)
	utils.CheckError(query.GetResults(&sourcesFromDB))

	if len(sourcesFromDB) != 0 {
		sourceFromReq.ID = sourcesFromDB[0].ID
		session.Load(&source, sourcesFromDB[0].ID)
		log.Info("Source ID: ", source.ID)
		source.Name = sourceFromReq.Name
		source.Address = sourceFromReq.Address
		source.Auth = sourceFromReq.Auth
		source.Type = sourceFromReq.Type
		source.Group = sourceFromReq.Group
		utils.CheckError(session.Store(source))
	} else {
		utils.CheckError(session.Store(sourceFromReq))
	}

	utils.CheckError(session.SaveChanges())
	ctx.JSON(http.StatusOK, sourceFromReq)
}

func Get(ctx *gin.Context) {
	var source ModuleSource = GetSource(ctx.Param("name"))
	log.Info("SOURCE ID: ", source.ID)

	if source.ID != "" {
		ctx.JSON(http.StatusOK, source)
	} else {
		ctx.String(http.StatusNotFound, fmt.Sprintf("Source %s not found", ctx.Param("name")))
	}
}

func GetSource(sourceName string) ModuleSource {
	var source *ModuleSource
	var sourcesFromDB []*ModuleSource
	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	query := session.QueryCollectionForType(reflect.TypeOf(&ModuleSource{})).WhereEquals("name", sourceName)
	utils.CheckError(query.GetResults(&sourcesFromDB))

	if len(sourcesFromDB) != 0 {
		utils.CheckError(session.Load(&source, sourcesFromDB[0].ID))
	} else {
		log.Warnf("Module source with name %s not found", sourceName)
		source = &ModuleSource{ID: ""}
	}

	return *source
}
