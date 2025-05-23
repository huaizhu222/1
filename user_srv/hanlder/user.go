package hanlder

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"user_srv/global"
	"user_srv/model"
	"user_srv/proto"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User_Server struct {
	proto.UnimplementedUserServer
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	// grpc的message中的字段用默认值，你不可以赋值为nil
	userInfoRsp := proto.UserInfoResponse{
		Id:            int32(user.User_id),
		Password:      user.Password,
		NickName:      user.NickName,
		Role:          uint32(user.Role),
		Like:          user.Like,
		LikeEmbedding: user.Embedding,
	}
	return userInfoRsp
}
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *User_Server) GetUserById(ctx context.Context, req *proto.UserInfoResponse) (*proto.UserInfoResponse, error) {
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
	user.Embedding = req.LikeEmbedding
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
