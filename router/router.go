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

	// 用户模块
	userBlock := r.Group("/user", middlewares.AuthCheck())
	{
		// 用户详情
		userBlock.GET("/detail", controllers.UserDetail())
	}

	return r
}
