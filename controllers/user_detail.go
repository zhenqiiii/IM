package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
)

// 获取用户详情
func UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "获取用户详情成功",
		})
	}
}
