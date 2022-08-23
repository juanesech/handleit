package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanesech/handleit/config"
	"github.com/juanesech/handleit/module"
)

var ListenAddr = "localhost:8080"

func main() {
	router := gin.Default()

	router.POST("/modules/import", module.Import)
	router.GET("/modules", module.List)
	router.GET("/modules/:name", module.Get)
	router.POST("/config", config.Set)

	router.Run()
}
