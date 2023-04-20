package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/user"
)

func Create(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong3",
		})
	})

	router.POST("/user", user.PostUser)
	// router.PUT("/somePut", user.GetUser)
	// router.DELETE("/someDelete", user.GetUser)
	router.GET("/user/:userid", user.GetUser)

}
