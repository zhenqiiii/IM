package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 获取用户详情(自己的个人)
func UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)

		// 查询用户数据
		user, err := sqldb.GetUserBasicByID(userInfo.UserID)
		if err != nil {
			log.Println("用户数据查询失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "数据查询异常：" + err.Error(),
			})
			return
		}

		// 获取成功，返回数据
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "获取用户详情成功",
			"data": user,
		})
	}
}
