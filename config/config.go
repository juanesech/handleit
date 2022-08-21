package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Conf Config

type Config struct {
	ModuleSource string `json:"moduleSource"`
}

func Set(ctx *gin.Context) {
	ctx.BindJSON(&Conf)
	ctx.JSON(http.StatusOK, Conf)
}
