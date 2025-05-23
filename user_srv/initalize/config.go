package initalize

import (
	"fmt"
	"user_srv/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool { //通过设置的环境变量，来改变配置文件yaml
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {

	v := viper.New()
	v.SetConfigFile("pro.yaml")
	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorw("连接yaml失败")
		panic(err)
	}
	// 取出yaml中的端口信息
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Errorw("取出yaml失败")
		panic(err)
	}
	// //从nacos中读取配置信息
	// serverConfigs := []constant.ServerConfig{
	// 	{
	// 		Port:   uint64(global.NacosConfig.Port),
	// 		IpAddr: global.NacosConfig.Host,
	// 	},
	// }
	// clientConfig := constant.ClientConfig{
	// 	NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	// 	TimeoutMs:           5000,
	// 	NotLoadCacheAtStart: true,
	// 	LogDir:              "tmp/nacos/log",
	// 	CacheDir:            "tmp/nacos/cache",
	// 	LogLevel:            "debug",
	// }
	// configClient, err := clients.CreateConfigClient(map[string]interface{}{
	// 	"serverConfigs": serverConfigs,
	// 	"clientConfig":  clientConfig,
	// })
	// if err != nil {
	// 	fmt.Printf("创建错误")
	// }

	// //从nacos中读取配置信息
	// content, err := configClient.GetConfig(vo.ConfigParam{
	// 	DataId: global.NacosConfig.DataId,
	// 	Group:  global.NacosConfig.Group,
	// })
	// if err != nil {
	// 	zap.S().Errorw("从nacos取出配置信息失败")
	// 	panic(err)
	// }

	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-srv.json",
	// 	Group:  "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("配置文件变化")
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })

	// if err != nil {
	// 	zap.S().Errorw("监视配置文件变化失败")
	// 	panic(err)
	// }

	// if err := json.Unmarshal([]byte(content), &global.ServerConfig); err != nil {
	// 	zap.S().Errorw("读取nacos配置失败%s", err)
	// }
	fmt.Println(global.ServerConfig)
}
