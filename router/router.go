package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/controllers"
	"github.com/zhenqiiii/IM-GO/middlewares"
)

// 路由创建
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置CORS，放行所有源
	r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000", "null"}, // 允许前端地址。允许 file:// 的 null 来源，因为可能会直接使用浏览器打开html
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
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
		userBlock.GET("/query", controllers.UserQuery())
		// 发送接收消息
		userBlock.GET("/msg", controllers.WebsocketMessage())
		// 拉取聊天室聊天记录列表(进入某房间时)
		userBlock.GET("/chatlist", controllers.ChatList())

		// 添加好友
		// 博主写的比较简单，但个人觉得一个完善的添加好友功能应该涉及发送请求，同意等过程
		userBlock.POST("/add", controllers.UserAdd())

		// 删除好友
		userBlock.DELETE("/delete", controllers.UserDelete())

	}

	return r
}
