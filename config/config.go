package config

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	ModuleSource string `json: "moduleSource"`
}

func Set(ctx *gin.Context) {
	var config Config
	ctx.BindJSON(&config)
	message := fmt.Sprintf("Module %s posted", config.ModuleSource)
	ctx.String(http.StatusOK, message)
}
