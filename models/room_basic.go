package models

import "time"

// 聊天室结构体
type RoomBasic struct {
	RoomID    string
	Name      string
	Info      string
	OwnerID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
