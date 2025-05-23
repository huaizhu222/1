package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeteled bool
}

type User struct {
	BaseModel
	User_id   uint
	Password  string `gorm:"type:varchar(100);not null comment '密码'"`
	NickName  string `gorm:"type:varchar(20) comment '用户名'"`
	Role      int    `gorm:"column:role;defualt:1;type:int comment '1表示普通用户,2表示管理员'"`
	Like      string `gorm:"type:varchar(20);not null comment '喜好'"`
	Embedding string `gorm:"type:varchar(20);"`
}
