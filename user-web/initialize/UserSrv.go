package initialize

import (
	"fmt"
	"log"
	"user-web/global"
	"user-web/proto"

	"google.golang.org/grpc"
)

func InitUserConn() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host, global.ServerConfig.UserSrvConfig.Port), // "127.0.0.1:50051"
		grpc.WithInsecure(), // 禁用 TLS（仅限测试环境）
	)
	if err != nil {
		log.Fatal(err)
	}
	userSrvClient := proto.NewUserClient(conn)
	global.UserSrvClient = userSrvClient
	if err != nil {
		log.Fatal(err)
	}
}
