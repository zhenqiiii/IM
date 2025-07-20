package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/redisdb"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/models"
	"github.com/zhenqiiii/IM-GO/pkg/snowflakeID"
	"golang.org/x/crypto/bcrypt"
)

// 用户注册路由处理函数：用户输入验证码后进行比对，如果对上了就注册成功，存入账号数据
// 参数：邮箱、账号、密码以及验证码
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
		code := c.PostForm("code")
		email := c.PostForm("email")
		account := c.PostForm("account")
		password := c.PostForm("password")
		// 参数校验
		// 1. 缺失校验
		if code == "" || email == "" || account == "" || password == "" {
			log.Println("注册参数缺失")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "参数缺失",
			})
			return
		}
		// 2. 判断账号account是否唯一
		exist, err := sqldb.CheckUserBasicExistByAccount(account)
		if err != nil {
			log.Println("sql查询失败")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统出错:" + err.Error(),
			})
			return
		}
		if exist {
			log.Println("用户输入账号已存在")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.ALREADY_EXISTS,
				"msg":  "该账号已被占用",
			})
			return
		}
		// 3. 验证码是否正确
		realCode, err := redisdb.Get(email)
		if err != nil {
			// 这里的error有两种情况:
			// 1. 查询过程中系统有问题报错
			// 2. 不存在该email对应的验证码而报Nil error,也就是用户还没有获取验证码
			// 一般是只有第2种情况的,所以这里给前端的响应只考虑第2种情况
			log.Println("从redis中获取验证码失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_STEPS,
				"msg":  "还未获取验证码：" + err.Error(),
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

		// 校验通过,存入用户数据
		// 预处理: 1. 使用bcrypt加密passoword然后存入 2. 使用雪花算法生成随机UserID
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := models.UserBasic{
			UserID:   snowflakeID.GenID(),
			Account:  account,
			Password: string(hashed),
			Email:    email,
		}
		// 插入数据库
		err = sqldb.InsertUserBasic(user)
		if err != nil {
			log.Println("创建用户失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}

		// 响应
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "注册成功",
		})
	}
}
