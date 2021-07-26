package controllers

import "github.com/gin-gonic/gin"

func ShowTeste(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
