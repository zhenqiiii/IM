package models

import (
	"time"
)

// 用户结构体
type UserBasic struct {
	UserID    string    `json:"userid" gorm:"primaryKey"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	Gender    int       `json:"gender"` // 0-未知 1-男 2-女
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
