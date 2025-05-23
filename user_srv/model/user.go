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
	Password  string `gorm:"type:varchar(100);not null "`
	NickName  string `gorm:"type:varchar(20) "`
	Role      int    `gorm:"column:role;defualt:1;type:int "`
	Like      string `gorm:"type:varchar(20);not null "`
	Embedding string `gorm:"type:varchar(20);"`
}
