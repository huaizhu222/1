package initialize

import (
	"fmt"
	"user-web/global"

	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool { //通过设置的环境变量，来改变配置文件yaml
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	configFileNamePrefix := "config_debug.yaml"
	v := viper.New()
	v.SetConfigFile(configFileNamePrefix)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 取出yaml中的端口信息
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	fmt.Println(global.ServerConfig)
}
