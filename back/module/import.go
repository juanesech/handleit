package module

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juanesech/topo/config"
	"github.com/juanesech/topo/constants"
	db "github.com/juanesech/topo/database"
	gl "github.com/juanesech/topo/gitlab"
	"github.com/juanesech/topo/utils"
)

func Import(ctx *gin.Context) {
	var rqConfig *ImportRequest
	var source config.ModuleSource
	var modsFromSource []*Module

	ctx.BindJSON(&rqConfig)

	source = config.GetSource(rqConfig.Name)

	switch ct := source.Type; ct {
	case "FileSystem":
		modsFromSource = getModulesFromFS(source.Address)

	case "GitLab":
		mp := gl.GetProjects(source, gl.GetGroup(source).Id)
		folder := uuid.NewString()

		for _, p := range mp {
			utils.Clone(fmt.Sprintf("%s/%s", folder, p.Name), source.Auth, p.Url)
		}
		modsFromSource = getModulesFromFS(fmt.Sprintf("/tmp/%s", folder))
	}

	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	for _, module := range modsFromSource {
		var modulesFromDB []*Module
		var loadedModule *Module

		query := session.QueryCollectionForType(reflect.TypeOf(&Module{})).WhereEquals("Name", module.Name)

		utils.CheckError(query.GetResults(&modulesFromDB))
		log.Info("START SAVE TO DB: ", module.Name)
		if len(modulesFromDB) != 0 {
			module.ID = modulesFromDB[0].ID
			session.Load(&loadedModule, module.ID)
			loadedModule.Variables = module.Variables
			loadedModule.Outputs = module.Outputs
			loadedModule.Providers = module.Providers
			log.Info("Document ID: ", loadedModule.ID)
			utils.CheckError(session.Store(loadedModule))
		} else {
			utils.CheckError(session.Store(module))
		}
		utils.CheckError(session.SaveChanges())
	}
}
