package router

import (
	"github.com/GuiCintra27/go/projects/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basepath := "/api/v1"

	handler.InitializeHandler()
	
	v1 := router.Group(basepath)
	{		
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.UpdateOpeningHandler)
		v1.PUT("/opening", handler.DeleteOpeningHandler)
	}
}