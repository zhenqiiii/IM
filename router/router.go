package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/controllers"
	"github.com/zhenqiiii/IM-GO/middlewares"
)

// 路由创建
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// router

	// login
	r.POST("/login", controllers.Login())
	r.GET("/index", controllers.GetIndex())
	// 注册时验证码发送
	r.POST("/verify", controllers.Send_Code())

	// 用户模块
	userBlock := r.Group("/user", middlewares.AuthCheck())
	{
		// 用户详情
		userBlock.GET("/detail", controllers.UserDetail())
		// 发送接收消息
		userBlock.GET("/msg", controllers.WebsocketMessage())
		// 拉取聊天室聊天记录列表(进入某房间时)
		userBlock.GET("/chatlist", controllers.ChatList())
	}

	return r
}
