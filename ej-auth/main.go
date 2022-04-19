package main

import (
	"github.com/aaraya0/arq-software/ejerciciosGo/ej-auth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login",
		controllers.Login)
	engine.Run(":8080")
}
