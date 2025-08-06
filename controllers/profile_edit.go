package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/models"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 编辑资料路由
func ProfileEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数
		var profile models.EditableProfileParams
		err := c.ShouldBind(&profile)
		if err != nil {
			log.Println("参数绑定错误：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "参数错误：" + err.Error(),
			})
		}

		// 获取用户标识
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		// 修改
		err = sqldb.UpdateProfile(userInfo.UserID, profile)
		if err != nil {
			log.Println("昵称修改失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常" + err.Error(),
			})
			return
		}

		// 成功
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "修改成功！",
		})

	}
}
