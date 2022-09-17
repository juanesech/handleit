package module

import (
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/juanesech/handleit/config"
	db "github.com/juanesech/handleit/database"
	"github.com/juanesech/handleit/utils"
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
		log.Info("Get from GitLab")
	}

	session, sessionErr := db.Client.OpenSession(db.Name)
	utils.CheckError(sessionErr)
	defer session.Close()

	for _, module := range modsFromSource {
		var modulesFromDB []*Module
		var loadedModule *Module

		query := session.QueryCollectionForType(reflect.TypeOf(&Module{})).WhereEquals("Name", module.Name)

		utils.CheckError(query.GetResults(&modulesFromDB))

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
