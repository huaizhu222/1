package initialize

import (
	"user-web/middlewares"
	"user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	g := gin.Default()
	//配置跨域
	g.Use(middlewares.Cors())
	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	ApiGroup := g.Group("/v1")
	{
		router.InitUserRouter(ApiGroup)
	}
	return g
}
