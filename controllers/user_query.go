package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 需要返回的用户信息
type userInfo struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"` // 0-未知 1-男 2-女
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"` // 是否为好友
}

// 查询指定用户的信息
func UserQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数；指定用户的account
		account := c.Query("account")
		if account == "" {
			log.Println("查询用户请求account参数为空")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "account参数缺失",
			})
			return
		}
		// 查询用户UserBasic
		user, err := sqldb.GetUserBasicByAccount(account)
		if err != nil {
			log.Println("查询失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "查询异常：" + err.Error(),
			})
			return
		}
		// 查询成功，返回结果:userInfo
		// 判断是否为好友
		err, friend := sqldb.JudgeTwoUsersAreFriends(c.MustGet("user_claims").(*jwt.UserClaims).UserID, user.UserID)
		if err != nil {
			log.Println("查询好友关系失败:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "查询异常:" + err.Error(),
			})
		}
		result := userInfo{
			Account:  user.Account,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Email:    user.Email,
			Avatar:   user.Avatar,
			IsFriend: friend,
		}
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "查询成功",
			"data": result,
		})
	}
}
