package controllers

import (
	"log"
	"net/http"

	"github.com/zhenqiiii/IM-GO/dao/redisdb"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/verification"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
)

// 此处的逻辑属于注册流程中的一部分（验证码发送）
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
		exist, err := sqldb.CheckUserBasicExistByEmail(email)
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

		// 参数处理通过，发送验证邮件
		code := verification.GenCode()
		err = verification.SendCode(email, code, "register") //register场景
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "验证码发送失败： " + err.Error(),
			})
			return
		}

		// 存储验证码用于比对(使用邮箱作为key)
		err = redisdb.Set(email, code)
		if err != nil {
			log.Println("验证码存储失败：" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统出错：" + err.Error(),
			})
			return
		}
		// 响应
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "验证码已发送!",
		})

	}
}
