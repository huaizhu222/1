package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int32  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int32  `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type ServerConfig struct {
	Name          string `mapstructure:"name" json:"name"`
	Port          int32  `mapstructure:"port" json:"port"`
	UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTConfig     `mapstructure:"jwt" json:"jwt"`
	RedisConfig   `mapstructure:"redis" json:"redis"`
}
