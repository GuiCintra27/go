package router

import (
	docs "github.com/GuiCintra27/go/projects/gopportunities/docs"
	"github.com/GuiCintra27/go/projects/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)
	{		
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}