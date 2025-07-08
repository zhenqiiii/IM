package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户结构体
type UserBasic struct {
	gorm.Model
	Name          string `gorm:" type:varchar(255)"`
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogOutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

// 为什么要定义这个方法？
func (table *UserBasic) TableName() string {
	return "user_basic"
}
