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
	// r.GET("/index", controllers.GetIndex())
	//用户注册
	r.POST("/register", controllers.Register())
	// 注册时验证码发送
	r.POST("/verify", controllers.Send_Code())

	// 用户模块
	userBlock := r.Group("/user", middlewares.AuthCheck())
	{
		// 用户详情（自己的信息）
		userBlock.GET("/detail", controllers.UserDetail())
		// 查看指定用户的个人信息
		userBlock.GET("/query", controllers.Query())
		// 发送接收消息
		userBlock.GET("/msg", controllers.WebsocketMessage())
		// 拉取聊天室聊天记录列表(进入某房间时)
		userBlock.GET("/chatlist", controllers.ChatList())
	}

	return r
}
