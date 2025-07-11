package controllers

import (
	"net/http"

	"github.com/zhenqiiii/IM-GO/gorm/sql"
	"github.com/zhenqiiii/IM-GO/pkg/verification"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
)

// 这里应该不是叫send_code,此处的逻辑应该是注册的逻辑
// 验证码发送处理函数:
// 场景：点击注册按钮后发送POST请求
// 接收参数：email
func Send_Code() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收邮箱参数
		email := c.PostForm("email")
		// 是否为空
		if email == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "邮箱不能为空",
			})
			return
		}
		// 是否已注册
		exist, err := sql.CheckUserBasicExistByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "查询数据失败：" + err.Error(),
			})
			return
		}
		if exist {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.ALREADY_EXISTS,
				"msg":  "用户已存在",
			})
			return
		}

		// 参数处理通过，发送验证码邮件
		code := verification.GenCode()
		err = verification.SendCode(email, code)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "验证码发送失败： " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "验证码已发送!",
		})

	}
}
