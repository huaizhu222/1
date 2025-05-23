package main

import (
	"context"
	"fmt"
	"log"
	"user_srv/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithInsecure(), // 禁用 TLS（仅限测试环境）
	)
	if err != nil {
		log.Fatal(err)
	}
	userSrvClient := proto.NewUserClient(conn)
	// userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
	// 	NickName: "小明",
	// 	PassWord: "123456789",
	// 	Like:     "篮球",
	// 	UserId:   1,
	// })

	// user, _ := userSrvClient.GetUserById(context.Background(), &proto.UserInfoResponse{
	// 	Id: 1,
	// })
	// fmt.Println(user)

	ok, _ := userSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:          "123456789",
		EncryptedPassword: "$pbkdf2-sha512$A7dgp5oywcqIEUA2$e6bbe530f4472b332dcb95611eeba664ea0c2a4afda69a25b8209f4f57451d20",
	})
	if ok.Success {
		fmt.Println(222222)
	}
}
