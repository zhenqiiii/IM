package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/shopping_system/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	// router
	r.GET("/index", controllers.GetIndex())

	return r
}
