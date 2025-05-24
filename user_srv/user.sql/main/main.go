package main

import (
	"fmt"
	"log"
	"os"
	"time"
	model "user_srv/user.sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// dsn := "root:123456@(127.0.0.1:3306)/user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=127.0.0.1 user=postgres password=123456 dbname=user_srv port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
	db.AutoMigrate(&model.User{})
}
