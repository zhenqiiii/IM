package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/gorm/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 登录处理函数
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数获取
		account := c.PostForm("account")
		password := c.PostForm("password")
		//参数缺失
		if account == "" || password == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "用户名或密码不能为空",
			})
			return
		}

		// 查询用户是否存在
		user, err := sqldb.GetUserBasicByAccount(account)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": cont.NOT_FOUND,
				"msg":  "查询失败：" + err.Error(),
			})
			return
		}
		// 存在，校验密码
		// todo:目标是用户输入明文，被前端哈希加密后传过来，然后和数据库中的哈希值比对
		// 目前存在数据库中的密码以及前端传过来的密码都是明文
		// 暂时将明文密码在此处加密再比较
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "密码错误：" + err.Error(),
			})
			return
		}

		// 校验通过，生成token
		token, err := jwt.GenToken(user.UserID, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "登录成功",
			"data": gin.H{
				"token": token,
			},
		})
	}
}
