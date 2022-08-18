package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var database []Module

type Module struct {
	Name      string     `json:"name"`
	Version   string     `json:"version"`
	Variables []variable `json:"variables"`
}

type variable struct {
	Name     string `json:"name"`
	VarType  string `json:"type"`
	Required bool   `json:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/modules", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, database)
	})

	router.POST("/modules", func(ctx *gin.Context) {
		var module Module
		ctx.BindJSON(&module)
		database = append(database, module)
		message := fmt.Sprintf("Module %s posted", module.Name)
		ctx.String(http.StatusOK, message)
	})

	router.Run()
}
