package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// index页路由处理（包含业务逻辑）
func GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "GetIndex",
		})
	}
}
