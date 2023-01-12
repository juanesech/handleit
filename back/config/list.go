package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
)

func List(ctx *gin.Context) {

    dbctx, dbclose := utils.GetCtx()
    defer dbclose()

    coll := db.GetCollection("sources")
    filter := bson.D{}
    var sourcesFromDB []ModuleSource
    cursor , finderr := coll.Find(dbctx, filter)
    utils.CheckError(finderr)
    utils.CheckError(cursor.All(dbctx, &sourcesFromDB))

	ctx.Header("Access-Control-Allow-Origin", "*")
//TODO: HTTP Response on error
	ctx.JSON(http.StatusOK, sourcesFromDB)
}
