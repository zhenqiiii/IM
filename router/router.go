package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/controllers"
)

// 路由创建
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// router
	r.GET("/index", controllers.GetIndex())

	return r
}
