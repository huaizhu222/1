package api

import (
	"context"
	"net/http"
	"time"
	"user-web/forms"
	"user-web/global"
	"user-web/middlewares"
	"user-web/models"
	"user-web/proto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func PassWordLogin(c *gin.Context) {

	PassWordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&PassWordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	//设置User_Id
	val, _ := c.Get("parentSpan")
	span, _ := val.(opentracing.Span)
	span.SetTag("user_id", PassWordLoginForm.User_Id)

	rsp, err := global.UserSrvClient.GetUserById(context.WithValue(context.Background(), "ginContext", c), &proto.UserInfoResponse{ // rsp就是用户信息
		Id: PassWordLoginForm.User_Id,
	})
	if err != nil {
		zap.S().Errorw("[PassWordLogin]查询 【用户密码失败】",
			"msg", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}
	ok, err := global.UserSrvClient.CheckPassword(context.WithValue(context.Background(), "ginContext", c), &proto.PasswordCheckInfo{
		EncryptedPassword: rsp.Password,
		Password:          PassWordLoginForm.PassWord,
	})

	if err != nil {
		zap.S().Errorw("[PassWordLogin]核对密码失败",
			"msg", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}

	if ok.Success { //密码正确
		//生成token
		j := middlewares.NewJWT()

		// ID 和NickName都是自己定义的Payload
		claims := models.CustomClaims{
			ID:          uint(rsp.Id),
			NickName:    rsp.NickName,
			AuthorityId: 1,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),               //签名的生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
				Issuer:    "imooc",                         //谁给的授权
			},
		}
		token, err := j.CreateToken(claims) //前两个部分head和Payload组成的claims和密钥j生成token
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成token失败",
			})
			zap.S().Errorw("生成token失败:", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":        "登录成功",
			"id":         rsp.Id,
			"nick_name":  rsp.NickName,
			"token":      token,
			"like":       rsp.Like,
			"expired_at": time.Now().Unix() + 60*60*24*30,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误",
		})
	}
}

func Register(c *gin.Context) {
	register := forms.RegisterForm{}
	// opation_id := time.Now().UnixNano()<<20 | (int64(1)&0x3FF)<<10

	if err := c.ShouldBind(&register); err != nil {
		HandleValidatorError(c, err)
		return
	}
	//设置tag
	val, _ := c.Get("parentSpan")
	span, _ := val.(opentracing.Span)
	span.SetTag("user_id", register.User_Id)

	//新建用户
	user, err := global.UserSrvClient.CreateUser(context.WithValue(context.Background(), "ginContext", c), &proto.CreateUserInfo{
		NickName: register.Nick_name,
		PassWord: register.PassWord,
		UserId:   register.User_Id,
		Like:     register.Like,
		// Mobile:   strconv.FormatInt(opation_id, 10),
	})

	if err != nil {
		zap.S().Errorw("新建用户失败", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	//新建用户成功，直接登录
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: 1, //用户为1，管理员为2
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24, //1天过期
		},
	}
	token, err := j.CreateToken(claims) //前两个部分head和Payload组成的claims和密钥j生成token
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		zap.S().Errorw("生成token失败:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":        "登录成功",
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func GetUser(c *gin.Context) {

	PassWordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&PassWordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	//设置user_Id
	val, _ := c.Get("parentSpan")
	span, _ := val.(opentracing.Span)
	span.SetTag("user_id", PassWordLoginForm.User_Id)

	rsp, err := global.UserSrvClient.GetUserById(context.WithValue(context.Background(), "ginContext", c), &proto.UserInfoResponse{ // rsp就是用户信息
		Id: PassWordLoginForm.User_Id,
	})
	if err != nil {
		zap.S().Errorw("[PassWordLogin]查询 【用户信息失败】",
			"msg", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":       "查询成功",
		"id":        rsp.Id,
		"nick_name": rsp.NickName,
		"like":      rsp.Like,
	})
}
