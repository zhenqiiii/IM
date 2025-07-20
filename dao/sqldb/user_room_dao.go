package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
)

// 存放user_room表的操作函数

// 通过UserID和RoomID查询user_room关系：即用户是否属于该聊天室
// 返回值：bool&error
func GetUserRoomByID(userid string, roomid string) (bool, error) {
	// 使用Find方法查询单个对象，同时调用Limit(1)方法速度会更快
	// 而且Find方法没有查询到记录时并不会报ErrRecordNotFound错误（First、Take、Last会有）
	result := db.Where("user_id = ? AND room_id = ? ", userid, roomid).Limit(1).Find(&models.UserRoom{})
	if result.Error != nil {
		log.Println("查询失败：" + result.Error.Error())
		return false, result.Error
	}
	// 查询到
	if result.RowsAffected != 0 {
		return true, nil
	}
	// 未查询到
	return false, nil
}

// 获取属于该房间的用户ID
// 返回值：UserRoom结构体切片
func GetUsersByRoomID(roomid string) (users []models.UserRoom, err error) {
	result := db.Find(&users)
	if result.Error != nil {
		log.Println("UserRoom查询失败：" + result.Error.Error())
		return nil, result.Error
	}
	return users, nil
}
