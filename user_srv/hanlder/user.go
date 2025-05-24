package hanlder

import (
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"strings"
	"user_srv/global"
	"user_srv/proto"
	model "user_srv/user.sql"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User_Server struct {
	proto.UnimplementedUserServer
}

func Float32SliceToString(embedding []float32) string {
	bytes, _ := json.Marshal(embedding)

	return string(bytes)
}

func StringToFloat32Slice(s string) []float32 {
	var embedding []float32
	json.Unmarshal([]byte(s), &embedding)
	return embedding
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:            int32(user.User_id),
		Password:      user.Password,
		NickName:      user.NickName,
		Role:          uint32(user.Role),
		Like:          user.Like,
		LikeEmbedding: Float32SliceToString(user.LikeEmbedding),
	}
	return userInfoRsp
}
func (s *User_Server) GetUserById(ctx context.Context, req *proto.UserInfoResponse) (*proto.UserInfoResponse, error) {
	// 开始链路追踪
	GetuserTracer := opentracing.SpanFromContext(ctx)
	// GetuserTracer := opentracing.GlobalTracer().StartSpan("GetUserbyId", opentracing.ChildOf(parentSpan.Context()))
	GetuserTracer.SetTag("user_id", req.Id)
	// 结束链路追踪
	defer GetuserTracer.Finish()

	var users model.User
	result := global.DB.Where(model.User{
		User_id: uint(req.Id),
	}).First(&users)

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInforsp := ModelToResponse(users)
	return &userInforsp, nil
}

func (s *User_Server) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) { //用户注册
	// 开始链路追踪
	parentSpan := opentracing.SpanFromContext(ctx)
	CreateUserTracer := opentracing.GlobalTracer().StartSpan("GetUserbyId", opentracing.ChildOf(parentSpan.Context()))
	CreateUserTracer.SetTag("user_id", req.UserId)
	defer CreateUserTracer.Finish()

	//新建用户
	var user model.User
	result := global.DB.Where("user_id = ?", req.UserId).First(&model.User{})
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在")
	}
	// 赋值
	user.User_id = uint(req.UserId)
	user.NickName = req.NickName
	user.Like = req.Like
	user.LikeEmbedding = StringToFloat32Slice(req.LikeEmbedding)
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//插入数据库
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *User_Server) CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckReponse, error) {
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckReponse{Success: check}, nil
}
