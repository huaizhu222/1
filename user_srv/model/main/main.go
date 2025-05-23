package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"user_srv/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func genMd5(code string) string { //用Md对密码进行加密
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}
func main() {
	dsn := "root:123456@(127.0.0.1:3306)/user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newlogger,
	})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
	db.AutoMigrate(&model.User{})

	// options := &password.Options{16, 100, 32, sha512.New}
	// salt, encodedPwd := password.Encode("admin123", options)
	// NewPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	// for i := 0; i < 10; i++ {
	// 	user := model.User{
	// 		NickName: fmt.Sprintf("bobby%d", i),
	// 		Password: NewPassword,
	// 	}
	// 	db.Save(&user)
	// }
}
