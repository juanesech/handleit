package module

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
    log "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/bson"
)

func Get(ctx *gin.Context) {
    module := Getmodule(ctx.Param("name"))
    log.Info("module ID: ", module.ID)

    if module.ID != "" {
        ctx.JSON(http.StatusOK, module)
    } else {
        ctx.String(http.StatusNotFound, fmt.Sprintf("Module %s not found", ctx.Param("name")))
    }
}

func Getmodule(moduleName string) Module {
    var module *Module
    dbctx, dbclose := utils.GetCtx()
    defer dbclose()

    coll := db.GetCollection("modules")
    filter := bson.D{{"name", moduleName}}
    findsrc := coll.FindOne(dbctx, filter).Decode(&module)
    utils.CheckError(findsrc)

    return module.WithID()
}
