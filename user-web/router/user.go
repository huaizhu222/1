package router

import (
	"user-web/api"
	"user-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Rounter *gin.RouterGroup) {
	UserRouter := Rounter.Group("/user").Use(middlewares.Trace()) // 使用链路追踪
	{
		UserRouter.POST("/login", api.PassWordLogin)              // 密码登录
		UserRouter.POST("/register", api.Register)                // 注册
		UserRouter.GET("/id", middlewares.JWTAuth(), api.GetUser) // jwt token验证 通过id查询用户信息
	}
}
