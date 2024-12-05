package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func handler(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"message": "ok"},
	)
}

func Run() {
	app := gin.Default()
	app.GET("/ping", handler)
	app.Run()
}
