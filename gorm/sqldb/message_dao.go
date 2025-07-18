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

// 获取消息列表
// 接收参数：房间id、分页长度limit、起始位置offset
func GetMessageListByRoomID(roomID string, limit int, offset int) ([]*models.MessageBasic, error) {
	list := make([]*models.MessageBasic, 0)
	// 查询数据库：消息内容按时间从晚到早排序（降序）即新到旧
	result := db.Where("room_id = ?", roomID).Order("created_at desc").Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		log.Println("访问消息表失败：" + result.Error.Error())
		return nil, result.Error
	}
	return list, nil
}
