package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 鉴权中间件：校验jwt并获取用户信息
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader("Authorization")

		// 解析token
		claims, err := jwt.ParseToken(token)
		if err != nil {
			// 中断处理
			c.Abort()
			log.Println("token认证不通过：" + err.Error())
			// 返回响应
			c.JSON(http.StatusBadRequest, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "用户认证不通过:" + err.Error(),
			})
			return
		}

		// token有效,将claims传入上下文供其他处理函数使用
		c.Set("user_claims", claims)
		// 执行下一个处理函数
		c.Next()
	}
}
