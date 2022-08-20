package config

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Conf Config

type Config struct {
	ModuleSource string `json:"moduleSource"`
}

func Set(ctx *gin.Context) {
	ctx.BindJSON(&Conf)
	message := fmt.Sprintf("Config %s posted", Conf.ModuleSource)
	ctx.String(http.StatusOK, message)
}
