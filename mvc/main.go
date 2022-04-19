package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":3000")
}


type Body struct {
	Name string `json:"name"`
}

func Test(ctx *gin.Context) {
	var body Body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(http.StatusOK, body)
}
func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")

}