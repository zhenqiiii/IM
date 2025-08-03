package models

import "time"

type MessageType string

const (
	MsgTypeText   MessageType = "text"
	MsgTypeEmoji  MessageType = "emoji"
	MsgTypeImage  MessageType = "image"
	MsgTypeFile   MessageType = "file"
	MsgTypeVideo  MessageType = "video"
	MsgTypeSpeech MessageType = "speech"
)

// 消息结构体
type MessageBasic struct {
	UserID    string      `json:"user_id" gorm:""`
	RoomID    string      `json:"room_id"`
	Data      string      `json:"data"`
	Type      MessageType `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// 指定数据库表名
func (MessageBasic) TableName() string {
	return "message_basic"
}
