package router

import (
	"user-web/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Rounter *gin.RouterGroup) {
	UserRouter := Rounter.Group("/user")
	{
		// UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("/login", api.PassWordLogin)
		UserRouter.POST("/register", api.Register)
	}
}
