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
	router.PUT("/user", user.PutUser)
	router.DELETE("/user/:userid", user.DelUser)
	router.GET("/user/:userid", user.GetUser)

}
