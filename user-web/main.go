package main

import (
	"fmt"
	"user-web/global"
	"user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	//初始化zap
	initialize.InitLogger()
	//初始化配置信息
	initialize.InitConfig()
	//初始化routers
	Router := initialize.Routers()
	//初始化grpc连接
	initialize.InitUserConn()

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
