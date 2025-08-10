package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/redisdb"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
)

// 重置路由，接收验证码比对并重置
func ResetPwd() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
		code := c.PostForm("code")
		email := c.PostForm("email")
		password := c.PostForm("password")
		// 参数校验
		// 缺失校验
		if code == "" || email == "" || password == "" {
			log.Println("Reset参数缺失")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "参数缺失",
			})
			return
		}
		// 验证码是否正确
		realCode, err := redisdb.Get(email)
		if err != nil {
			log.Println("从redis中获取验证码失败:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_STEPS,
				"msg":  "还未获取验证码:" + err.Error(),
			})
			return
		}
		// 比对
		if realCode != code {
			log.Println("用户输入验证码不正确")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "验证码错误",
			})
			return
		}

		// Reset Pwd
		// 获取UserId
		user, err := sqldb.GetUserBasicByEmail(email)
		if err != nil {
			log.Println("获取id失败:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}
		err = sqldb.UpdatePwd(password, user.UserID)
		if err != nil {
			log.Println("更新密码失败:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}

		// 响应
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "密码重置成功",
		})
	}
}
