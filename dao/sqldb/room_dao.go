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

// 删除房间
func DeleteRoomBasic(roomid string) error {
	result := db.Where("room_id = ?", roomid).Delete(&models.RoomBasic{})
	if result.Error != nil {
		log.Println("删除房间数据失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}

// 通过RoomID获取RoomBasic
func GetRoomBasicByRoomID(id string) (room *models.RoomBasic, err error) {
	result := db.Where("room_id = ?", id).Limit(1).Find(&room)
	if result.Error != nil {
		log.Println("[DB]GetRoomBasicByRoomID:" + result.Error.Error())
		return nil, result.Error
	}
	return room, nil
}
