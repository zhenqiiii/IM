package models

import "time"

// 消息结构体
type MessageBasic struct {
	UserID    string    `json:"userid" gorm:""`
	RoomID    string    `json:"roomid"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 指定数据库表名
func (MessageBasic) TableName() string {
	return "message_basic"
}
