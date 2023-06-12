package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/jnhu76/dwz/docs"
	"github.com/jnhu76/dwz/routers/api"
	v1 "github.com/jnhu76/dwz/routers/api/v1"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.GetAuth)
	r.GET("/hello", api.GetHello)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/add", v1.AddUrl)
	apiv1.GET("/:shorten", v1.GetUrl)
	apiv1.DELETE("/:shorten", v1.DeleteUrl)
	// apiv1.Use(jwt.JWT())
	// {
	// 	// test
	// 	apiv1.GET("/jwt", v1.GetJwt)
	// }

	return r
}
