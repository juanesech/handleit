package module

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juanesech/topo/config"
	db "github.com/juanesech/topo/database"
	gl "github.com/juanesech/topo/gitlab"
	"github.com/juanesech/topo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Import(ctx *gin.Context) {
	var rqConfig *ImportRequest
	var source config.ModuleSource
	var modsFromSource []Module
	dbctx, dbclose := utils.GetCtx()
	defer dbclose()

	utils.CheckError(ctx.BindJSON(&rqConfig))

	source = config.GetSource(rqConfig.Name)

	switch st := source.Type; st {
	case "FileSystem":
		modsFromSource = getModulesFromFS(source.Address)

	case "GitLab":
		mp := gl.GetProjects(source, source.Group)
		folder := uuid.NewString()

		for _, p := range mp {
			utils.Clone(fmt.Sprintf("%s/%s", folder, p.Name), source.Auth, p.Url)
		}
		if len(mp) != 0 {
			modsFromSource = getModulesFromFS(fmt.Sprintf("/tmp/%s", folder))
		}
	}

	coll := db.GetCollection("modules")

	for _, m := range modsFromSource {
		log.Info("IMPORTING: ", m.Name)
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{"name", m.Name}}
		_, upderr := coll.UpdateOne(dbctx, filter, bson.D{{"$set", m}}, opts)
		utils.CheckError(upderr)
	}
}
