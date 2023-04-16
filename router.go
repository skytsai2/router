package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong3",
		})
	})

}
