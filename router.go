package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/controller"
	"github.com/skytsai2/middleware"
)

type Data struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
	Users     []User `json:"data"`
}

type User struct {
	ID   int
	Name string
	Tel  string
}

func Create(router *gin.Engine) {

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {

		bodystr := sendHttpRequest("http://localhost:8080/api/user/list")

		var data Data
		err := json.Unmarshal([]byte(bodystr), &data)
		fmt.Println(bodystr)
		fmt.Println(data)
		fmt.Println(err)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"users": data.Users,
		})
	})

	router.GET("/test", func(c *gin.Context) {

		bodystr := sendHttpRequest("http://localhost:8080/api/user/list")

		var data Data
		err := json.Unmarshal([]byte(bodystr), &data)
		fmt.Println(bodystr)
		fmt.Printf("%s, %s", data.ErrorCode, data.Status)
		fmt.Println(err)
	})

	api := router.Group("/api")
	api.Use(middleware.Auth())
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/user", controller.PostUser)
		api.PUT("/user", controller.PutUser)
		api.DELETE("/user/:userid", controller.DelUser)
		api.GET("/user/:userid", controller.GetUser)
		api.GET("/user/list", controller.GetUserList)
	}
}

func sendHttpRequest(url string) string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Add("Key", "123")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	var bodystr string
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		bodystr = string(body)
		if err != nil {
			log.Fatal(err)
			return ""
		}
	} else {
		return ""
	}

	return bodystr
}
