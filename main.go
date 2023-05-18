package main

import (
	"log"
	"time"
	"url/models"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	store := persistence.NewInMemoryStore(5 * 60 * time.Second)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	router.POST("/new", models.PostUrl) // url
	// router.GET("/:url", models.GetUrl)
	router.GET("/:url", cache.CachePage(store, 60*time.Minute, models.GetUrl))
	router.DELETE("/:url", models.DeleteUrl)

	return router
}

func main() {
	router := setupRouter()

	models.ConnectDatabase()

	log.Fatal(router.Run("127.0.0.1:9090"))
}
