package config

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Config struct {
	Sources []ModuleSource
}

type ModuleSource struct {
	ID      string
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Group   int    `json:"group"`
	Auth    string `json:"auth"`
}

func (m ModuleSource) WithID() ModuleSource {
	var sourcefromdb bson.M

	dbctx, dbclose := utils.GetCtx()
	defer dbclose()

	coll := db.GetCollection("sources")
	filter := bson.D{{"name", m.Name}}
	findsrc := coll.FindOne(dbctx, filter).Decode(&sourcefromdb)
	utils.CheckError(findsrc)
	m.ID = sourcefromdb["_id"].(primitive.ObjectID).Hex()

	return m
}

func Set(ctx *gin.Context) {
	var sourceFromReq *ModuleSource

	utils.CheckError(ctx.BindJSON(&sourceFromReq))

	dbctx, dbclose := utils.GetCtx()
	defer dbclose()

	coll := db.GetCollection("sources")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"name", sourceFromReq.Name}}
	_, upderr := coll.UpdateOne(dbctx, filter, bson.D{{"$set", sourceFromReq}}, opts)
	utils.CheckError(upderr)

	ctx.JSON(http.StatusOK, sourceFromReq.WithID())
}

func Get(ctx *gin.Context) {
	source := GetSource(ctx.Param("name"))
	log.Info("SOURCE ID: ", source.ID)

	if source.ID != "" {
		ctx.JSON(http.StatusOK, source)
	} else {
		ctx.String(http.StatusNotFound, fmt.Sprintf("Source %s not found", ctx.Param("name")))
	}
}

func GetSource(sourceName string) ModuleSource {
	var source *ModuleSource
	dbctx, dbclose := utils.GetCtx()
	defer dbclose()

	coll := db.GetCollection("sources")
	filter := bson.D{{"name", sourceName}}
	findsrc := coll.FindOne(dbctx, filter).Decode(&source)
	utils.CheckError(findsrc)

	return source.WithID()
}
