package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanesech/handleit/config"
	"github.com/juanesech/handleit/module"
)

func main() {
	router := gin.Default()

	router.GET("/modules", module.List)
	router.GET("/modules/:name", module.Get)
	router.POST("/modules", module.New)
	router.POST("/config", config.Set)

	router.Run()
}
