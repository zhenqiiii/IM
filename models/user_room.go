package models

import "time"

// 私聊可以理解为一个只有我和对方两人的聊天室

// 用户与聊天室的关联表
type UserRoom struct {
	UserID    string
	RoomID    string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (UserRoom) TableName() string {
	return "user_room"
}
