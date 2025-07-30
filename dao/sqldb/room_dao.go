package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
)

// 插入房间
func InsertRoomBasic(room models.RoomBasic) error {
	result := db.Create(&room)
	if result.Error != nil {
		log.Println("插入房间数据失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}
