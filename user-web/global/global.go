package global

import (
	"user-web/config"
	"user-web/proto"


)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient
)
