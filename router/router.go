package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/controllers"
	"github.com/zhenqiiii/IM-GO/middlewares"
)

// 路由创建
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// // 把 ./frontend 目录下的所有文件托管到/static路径
	r.Static("/static", "./frontend")

	// // 若前后端在本机不同端口运行，配置CORS，放行所有源
	// r.Use(cors.Default())

	// 个性化配置CORS
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
	r.POST("/verify", controllers.SendCode())

	// 已登录
	userBlock := r.Group("/user", middlewares.AuthCheck())
	{
		// 用户详情（自己的信息）
		userBlock.GET("/detail", controllers.UserDetail())
		// 资料编辑
		userBlock.POST("/edit", controllers.ProfileEdit())
		// 密码修改
		userBlock.POST("/pwdchange", controllers.PwdChange())
		// 查看指定用户的个人信息
		userBlock.GET("/query", controllers.UserQuery())
		// 发送接收消息
		userBlock.GET("/msg", controllers.WebsocketMessage())
		// 拉取聊天室聊天记录列表(进入某房间时)
		userBlock.GET("/history", controllers.MsgHistory())
		// 获取联系人列表
		userBlock.GET("/chatlist", controllers.ChatList())

		// 添加好友
		// TODO：完善同意流程
		userBlock.POST("/add", controllers.UserAdd())

		// 删除好友
		userBlock.DELETE("/delete", controllers.UserDelete())

	}

	return r
}
