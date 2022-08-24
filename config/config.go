package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juanesech/handleit/database"
	"github.com/juanesech/handleit/utils"
)

type Config struct {
	ModuleSource string `json:"moduleSource"`
}

func Set(ctx *gin.Context) {
	var conf *Config

	session, sessionErr := db.Client.OpenSession("handleit")
	utils.CheckError(sessionErr)
	defer session.Close()

	utils.CheckError(session.Load(&conf, "configs/33-A"))
	utils.CheckError(session.Store(conf))
	utils.CheckError(session.SaveChanges())

	ctx.JSON(http.StatusOK, conf)
}

func Get() *Config {
	var conf *Config
	session, sessionErr := db.Client.OpenSession("handleit")
	utils.CheckError(sessionErr)
	defer session.Close()

	utils.CheckError(session.Load(&conf, "configs/33-A"))

	return conf
}
