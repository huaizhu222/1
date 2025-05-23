package initalize

import (
	"fmt"
	"log"
	"os"
	"time"
	"user_srv/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	c := global.ServerConfig.PgsqlConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.User, c.Password, c.Name, c.Port)
	fmt.Println(dsn)
	// 连接数据库
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newlogger,
	})
	global.DB = db
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
}
