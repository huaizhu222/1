package global

import (
	"user-web/config"
	"user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
)
