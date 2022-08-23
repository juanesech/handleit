package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juanesech/handleit/database"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ModuleSource string `json:"moduleSource"`
}

func Set(ctx *gin.Context) {
	var conf *Config
	session, err := db.Client.OpenSession("handleit")
	if err != nil {
		log.Warn("Failed to load config", err)
		ctx.BindJSON(&conf)
	}
	err = session.Load(&conf, "configs/33-A")
	if err != nil {
		log.Fatalf("session.Load() failed with %s\n", err)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		log.Fatalf("Can't open database session ", err)
	}

	err = session.Store(conf)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		log.Error("Can't open database session ", err)
	}
	err = session.SaveChanges()
	if err != nil {
		log.Error("Save failed ", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if err == nil {
		ctx.JSON(http.StatusOK, conf)
	}
	ctx.JSON(http.StatusOK, conf)
}

func Get() *Config {
	var conf *Config
	session, err := db.Client.OpenSession("handleit")
	if err != nil {
		log.Error("Failed to load config", err)
	}
	err = session.Load(&conf, "configs/33-A")
	if err != nil {
		log.Error("Failed to load config", err)
	}

	return conf
}
