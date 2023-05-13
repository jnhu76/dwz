package main

import (
	"log"
	"url/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	router.POST("/new", models.PostUrl) // url
	router.GET("/:url", models.GetUrl)
	router.DELETE("/:url", models.DeleteUrl)

	return router
}

func main() {
	router := setupRouter()

	models.ConnectDatabase()

	log.Fatal(router.Run("0.0.0.0:9090"))
}
