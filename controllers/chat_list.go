package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 联系人结构体
// Name: 名称
// RoomID: 所属房间id,用于后续获取聊天记录
// Type: 类型 1私聊 2群聊, -1表示数据获取出错情况，用于在聊天人列表中显示聊天类型
type Contact struct {
	Name   string `json:"name"`
	RoomID string `json:"room_id"`
	Type   int    `json:"type"`
}

// 获取联系人（好友&群聊）列表
func ChatList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		// 查询联系人列表
		// 执行逻辑：查询该用户的所有UserRoom记录，type=1的记录GetAnotherUser，type=0的获取房间名
		URList, err := sqldb.GetURListByUserID(userInfo.UserID)
		if err != nil {
			log.Println("获取联系人列表失败" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		// 对URList进行处理，获得contactList联系人列表
		// TODO:如果其中有一个用户的信息获取出现错误，如何处理？
		// 当前处理逻辑：出错后对该条数据另设内容，以表示错误情况，但依旧放在列表中
		// Debug: 获取的用户Nickname为空
		contactList := make([]*Contact, len(URList))
		for cnt, ur := range URList {
			contactList[cnt] = &Contact{} // 初始化slice
			switch ur.RoomType {
			case 1: // 好友（私聊）
				id, err := sqldb.GetAnotherUserID(ur)
				if err != nil {
					log.Println("获取另一用户id失败：" + err.Error())
					contactList[cnt].Name = "获取错误，刷新以重试"
					contactList[cnt].RoomID = ""
					contactList[cnt].Type = -1
					continue
				}
				contact, err := sqldb.GetUserBasicByID(id)
				if err != nil {
					log.Println("获取另一用户信息失败：" + err.Error())
					contactList[cnt].Name = "获取错误，刷新以重试"
					contactList[cnt].RoomID = ""
					contactList[cnt].Type = -1
					continue
				}
				contactList[cnt].Name = contact.Nickname
				contactList[cnt].RoomID = ur.RoomID
				contactList[cnt].Type = 1
			case 0: //群聊
				room, err := sqldb.GetRoomBasicByRoomID(ur.RoomID)
				if err != nil {
					log.Println("获取群聊房间信息失败：" + err.Error())
					// c.JSON(http.StatusOK, gin.H{
					// 	"code": cont.INTERNAL_ERROR,
					// 	"msg":  "系统异常：" + err.Error(),
					// })
					contactList[cnt].Name = "获取错误，刷新以重试"
					contactList[cnt].RoomID = ""
					contactList[cnt].Type = -1
					continue
				}
				contactList[cnt].Name = room.Name
				contactList[cnt].RoomID = room.RoomID
				contactList[cnt].Type = 0
			}
		}
		// 返回联系人列表
		// 暂时假设都会获取成功
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "联系人列表获取成功",
			"data": gin.H{
				"chatlist": contactList,
			},
		})

	}
}
