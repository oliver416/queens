package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	docs "queens/app/docs"
)

// Ping godoc
// @Summary Ping
// @Schemes
// @Description Test endpoint
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
func handler(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"message": "ok"},
	)
}

func Run() {
	app := gin.Default()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	group := app.Group(basePath)
	{
		group.GET("/ping", handler)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run()
}
