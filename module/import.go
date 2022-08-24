package module

import (
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	db "github.com/juanesech/handleit/database"
	"github.com/juanesech/handleit/utils"
)

func Import(ctx *gin.Context) {
	var modulesFromFS []*Module = getModulesFromFS()

	session, sessionErr := db.Client.OpenSession("handleit")
	utils.CheckError(sessionErr)
	defer session.Close()

	for _, module := range modulesFromFS {
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
