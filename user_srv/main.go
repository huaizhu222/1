package main

import (
	"flag"
	"fmt"
	"net"
	"user_srv/hanlder"
	"user_srv/initalize"
	"user_srv/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//初始化
	initalize.InitLogger()
	initalize.InitConfig()
	initalize.InitDB()

	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &hanlder.User_Server{})
	proto.RegisterSystemServer(server, &hanlder.Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		zap.S().Errorw("连接失败")
		return
	}
	err = server.Serve(lis) //启动端口
	if err != nil {
		fmt.Println("错误:", err)
	}
}
