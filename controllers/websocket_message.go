package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// author: zhenqiiii
// 理解：
// 单个用户和服务端只建立一个websocket连接，
// 用户在不同聊天室发送的消息都通过该WebSocket连接传输，
// 不同聊天室的消息根据消息体MessageBody中的RoomID字段区分

// 用于存放客户端发送的消息内容，用户身份从上下文中获得
type MessageBody struct {
	Message string `json:"message"`
	RoomID  string `json:"roomid"` //区分聊天室
}

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

		// 循环：持续读取消息内容
		// 用户自己在客户端发送的消息在此接收，然后分发给所有人
		for {
			msg := new(MessageBody)
			err = conn.ReadJSON(msg)
			// fmt.Println(msg.Message)
			if err != nil {
				log.Println("消息接收失败：" + err.Error())
				// c.JSON(http.StatusOK, gin.H{
				// 	"code": cont.INTERNAL_ERROR,
				// 	"msg":  "消息发送失败：" + err.Error(),
				// })
				return
			}
			// 接收成功,发送给所有人(同一聊天室)
			for _, cc := range conns {
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
