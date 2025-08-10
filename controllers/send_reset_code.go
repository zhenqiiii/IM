package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/redisdb"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/verification"
)

// 发送忘记密码后的密码重置验证邮件：需要确定用户输入的邮箱存在账号
func SendResetCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		// email参数
		email := c.PostForm("email")
		if email == "" {
			log.Println("email为空")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "缺少邮箱",
			})
			return
		}
		// 查询账号是否存在
		exist, err := sqldb.CheckUserBasicExistByEmail(email)
		if err != nil {
			log.Println("查询出错:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		// 不存在
		if !exist {
			log.Println("该邮箱还未注册账号")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.NOT_FOUND,
				"msg":  "该账号不存在，请先注册",
			})
			return
		}

		// 发送验证邮件
		code := verification.GenCode()
		err = verification.SendCode(email, code, verification.ResetMode) //忘记密码重置场景
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
			c.JSON(http.StatusOK, gin.H{
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
