package initalize

import "go.uber.org/zap"

func InitLogger() {
	// l := zap.NewDevelopmentConfig()
	// l.OutputPaths = []string{
	// 	"stdout",
	// }
	// logger, _ := l.Build()
	//更换全局日志
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

}
