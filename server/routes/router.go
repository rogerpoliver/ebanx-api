package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogerpoliver/ebanx-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		main.GET("/balance", controllers.GetBalance)
		main.POST("/event", controllers.PostEvent)
		main.POST("/reset", controllers.ResetData)
	}
	return router
}
