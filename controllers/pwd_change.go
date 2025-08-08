package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 修改密码
// 接收参数：raw旧密码、raw新密码--form表单
func PwdChange() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数
		oldPwd := c.PostForm("oldpwd")
		newPwd := c.PostForm("newpwd")
		if oldPwd == "" || newPwd == "" {
			log.Println("修改密码参数缺失")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "缺少参数",
			})
			return
		}
		// 用户标识
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		// 新密码比对由前端完成
		//校验旧密码
		correct, err := sqldb.VerifyPwd(oldPwd, userInfo.UserID)
		if err != nil {
			log.Println("校验密码出错：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		if !correct {
			log.Println("用户原密码输入错误")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "原密码输入错误",
			})
			return
		}
		err = sqldb.UpdatePwd(newPwd, userInfo.UserID)
		if err != nil {
			log.Println("修改密码出错：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "修改成功！",
		})

	}
}
