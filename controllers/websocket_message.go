package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/models"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// author: zhenqiiii
// 理解：
// 单个用户和服务端只建立一个websocket连接，
// 用户在不同聊天室发送的消息都通过该WebSocket连接传输，
// 不同聊天室的消息根据消息体MessageBody中的RoomID字段区分

// 用于存放客户端发送的消息内容，用户身份从上下文中获得
// type MessageBody struct {
// 	Message string `json:"message"`
// 	RoomID  string `json:"room_id"` //区分聊天室
// }

// websocket升级实例
var upgrader = websocket.Upgrader{}

// websocket连接池
var conns = make(map[string]*websocket.Conn)

// 消息发送接收：基于WebSocket
func WebsocketMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 升级连接到WebSocket
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("连接升级失败：" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "websocket连接失败：" + err.Error(),
			})
			return
		}
		defer conn.Close()

		// 从上下文中获得用户身份,并存储连接
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		conns[userInfo.UserID] = conn

		// 循环：保持连接，读取消息内容
		// 用户自己在客户端发送的消息在此接收，然后分发给其他用户
		// TODO：用户刷新页面导致的连接中断处理-心跳检测？
		for {
			// 获取消息
			msg := new(models.MessageBasic)
			err = conn.ReadJSON(msg)
			if err != nil {
				log.Println("消息接收失败：" + err.Error())
				log.Println("用户可能刷新页面导致websocket连接中断")
				return
			}
			// 判断用户是否属于消息体中RoomID字段对应的房间
			// 使用参数：UserID、RoomID	对user_room表进行查询
			belong, err := sqldb.GetUserRoomByID(userInfo.UserID, msg.RoomID)
			if err != nil {
				log.Println("用户-房间关系查询失败：" + err.Error())
				return
			}
			if !belong {
				log.Println("用户不属于该房间")
				return
			}
			// 保存消息
			err = sqldb.InsertMessageBasic(*msg)
			if err != nil {
				log.Println("消息保存失败：" + err.Error())
				return
			}

			// 获取属于该RoomID房间的所有用户
			users, err := sqldb.GetUsersByRoomID(msg.RoomID)
			if err != nil {
				log.Println("获取房间用户失败：" + err.Error())
				return
			}
			// 用户列表获取成功，遍历用户列表发送消息
			// 这里通过判断连接池中是否有对应用户来实现将消息只发送给在线用户
			// 离线用户通过获取消息列表来查看离线时聊天室已发送的消息
			for _, user := range users {
				if cc, online := conns[user.UserID]; online {
					// err = cc.WriteMessage(websocket.TextMessage, []byte(msg.Message))
					err = cc.WriteJSON(msg)
					if err != nil {
						log.Println("消息转发失败：" + err.Error())
						return
					}
				}
			}
		}

	}
}
