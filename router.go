package router

import (
	"fmt"
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

func test() {
	fmt.Println("124")
}
