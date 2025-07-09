package models

import "time"

// 用户与聊天室的关联表
type User_Room struct {
	UserID    string
	RoomID    string
	MessageID string
	CreatedAt time.Time
	UpdateAt  time.Time
}
