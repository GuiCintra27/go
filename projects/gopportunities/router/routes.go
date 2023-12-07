package router

import (
	docs "github.com/GuiCintra27/go/projects/gopportunities/docs"
	"github.com/GuiCintra27/go/projects/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basepath := "/api/v1"

	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = basepath
	
	v1 := router.Group(basepath)
	{		
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}