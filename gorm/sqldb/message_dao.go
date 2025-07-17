package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
)

// 存放Message_Basic数据库表相关方

// 插入消息
func InsertMessageBasic(msg models.MessageBasic) error {
	result := db.Create(&msg)
	if result.Error != nil {
		log.Println("消息创建失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}
