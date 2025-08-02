package sqldb

import (
	"errors"
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

// 根据A用户的UserRoom查询私聊房间中的B用户id
func GetAnotherUserID(ur *models.UserRoom) (string, error) {
	users, err := GetUsersByRoomID(ur.RoomID)
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if user.UserID != ur.UserID {
			return user.UserID, nil
		}
	}
	// users中未找到另一用户
	log.Println("[DB]GetAnotherUser error")
	return "", errors.New("[DB]GetAnotherUser")
}

// 查询二者是否为好友：
// 接收参数：id1, id2
// 查询二者之一的UserRoom表中RoomType为1的记录，取room_id字段,得到列表
// 然后用列表去匹配id2的UserRoom表记录
func JudgeTwoUsersAreFriends(id1, id2 string) (error, bool) {
	// 首先查询id1的所有Roomtype为1（私聊）的UserRoom记录
	// 得到的结果实际上就是该用户的所有私聊房间（即私聊好友列表）
	friendList := make([]models.UserRoom, 0)
	result := db.Where("user_id = ? AND room_type = 1", id1).Find(&friendList)
	if result.Error != nil {
		log.Println("获取私聊房间列表失败" + result.Error.Error())
		return result.Error, false
	}
	// 取出RoomID
	// TODO:优化此处的执行逻辑
	roomIdList := make([]string, 0)
	for _, id := range friendList {
		roomIdList = append(roomIdList, id.RoomID)
	}
	// 拿到id1的私聊列表后,查询id2的UserRoom表中是否存在一条记录,
	// 该记录满足条件:room_id in roomIdList
	result = db.Where("user_id = ? AND room_id IN ?", id2, roomIdList).Limit(1).Find(&models.UserRoom{})
	if result.Error != nil {
		log.Println("查询是否匹配失败" + result.Error.Error())
		return result.Error, false
	}
	// 查询成功
	// 不是好友关系
	if result.RowsAffected == 0 {
		return nil, false
	}
	// 是
	return nil, true
}

// 查找二者的私聊房间ID
func GetTwoUsersRoom(id1, id2 string) (string, error) {
	// 首先查询id1的所有Roomtype为1（私聊）的UserRoom记录
	// 得到的结果实际上就是该用户的所有私聊房间（即私聊好友列表）
	friendList := make([]models.UserRoom, 0)
	result := db.Where("user_id = ? AND room_type = 1", id1).Find(&friendList)
	if result.Error != nil {
		log.Println("获取私聊房间列表失败" + result.Error.Error())
		return "", result.Error
	}
	// 取出RoomID
	// TODO:优化此处的执行逻辑
	roomIdList := make([]string, 0)
	for _, id := range friendList {
		roomIdList = append(roomIdList, id.RoomID)
	}
	// 拿到id1的私聊列表后,查询id2的UserRoom表中是否存在一条记录,
	// 该记录满足条件:room_id in roomIdList
	userRoom := models.UserRoom{}
	result = db.Where("user_id = ? AND room_id IN ?", id2, roomIdList).Limit(1).First(&userRoom)
	if result.Error != nil {
		log.Println("查询房间失败" + result.Error.Error())
		return "", result.Error
	}
	// 查询成功,返回RoomID
	return userRoom.RoomID, nil
}

// 通过用户ID获取UserRoom列表:用于获取用户联系人列表
func GetURListByUserID(id string) (URList []*models.UserRoom, err error) {
	result := db.Where("user_id = ?", id).Find(&URList)
	if result.Error != nil {
		log.Println("[DB]GetURListByUserID：" + result.Error.Error())
		return URList, result.Error
	}
	return URList, nil

}

// 插入UserRoom关系
func InsertUserRoom(userRoom *models.UserRoom) error {
	result := db.Create(&userRoom)
	if result.Error != nil {
		log.Println("插入UserRoom用户房间关系数据失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}

// 删除私聊UserRoom关系
func DeleteUserRoom(roomid string) error {
	result := db.Where("room_id = ?", roomid).Delete(&models.UserRoom{})
	if result.Error != nil {
		log.Println("删除UserRoom用户房间关系数据失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}
