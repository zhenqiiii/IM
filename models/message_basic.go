package models

import "time"

// 消息结构体
type MessageBasic struct {
	UserID    string
	RoomID    string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
