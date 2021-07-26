package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogerpoliver/ebanx-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		teste := main.Group("teste")
		{
			teste.GET("/", controllers.ShowTeste)
		}
	}
	return router
}
