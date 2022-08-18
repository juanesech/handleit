package module

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Just for testing proposes
var Database []Module

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

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Database)
}

func New(ctx *gin.Context) {
	var newModule Module
	ctx.BindJSON(&newModule)
	Database = append(Database, newModule)
	message := fmt.Sprintf("Module %s posted", newModule.Name)
	ctx.String(http.StatusOK, message)
}
