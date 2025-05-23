package router

import (
	"user-web/api"
	"user-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Rounter *gin.RouterGroup) {
	UserRouter := Rounter.Group("/user")
	{
		// UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("/login", api.PassWordLogin)
		UserRouter.POST("/register", api.Register)
		UserRouter.GET("/id", middlewares.JWTAuth(), api.GetUser)
	}
}
