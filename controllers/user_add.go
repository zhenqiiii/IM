package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/models"
	"github.com/zhenqiiii/IM-GO/pkg/genid"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 用户添加路由（简化版）
// 接收参数：指定用户账号account：form表单
func UserAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收参数
		account := c.PostForm("account")
		if account == "" {
			log.Println("请求携带的account参数为空")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "账号参数为空",
			})
			return
		}

		// 查询被添加用户UserBasic
		user, err := sqldb.GetUserBasicByAccount(account)
		if err != nil {
			log.Println("查询失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}

		// 判断用户和被添加用户是否已经是好友关系
		// TODO：这里应该有bug
		uClaims := c.MustGet("user_claims").(*jwt.UserClaims)
		err, friend := sqldb.JudgeTwoUsersAreFriends(uClaims.UserID, user.UserID)
		if err != nil {
			log.Println("查询好友关系失败:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常:" + err.Error(),
			})
			return
		}
		// 已为好友
		if friend {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.ALREADY_EXISTS,
				"msg":  "已经是好友，不可重复添加",
			})
			return
		}
		// 不是好友，创建好友关系（创建私聊房间）
		// 创建三条表项：RoomBasic、A用户的UserRoom、B用户的UserRoom
		// 1. RoomBasic
		room := models.RoomBasic{
			RoomID:  genid.GenRoomID(),
			Name:    " 私聊",
			Info:    " 用户A(owner):" + uClaims.UserID + "\n用户B:" + user.UserID,
			OwnerID: uClaims.UserID,
		}
		err = sqldb.InsertRoomBasic(room)
		if err != nil {
			log.Println("创建私聊失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		// 2. UserRoom
		AuserRoom := &models.UserRoom{
			UserID:   uClaims.UserID,
			RoomID:   room.RoomID,
			RoomType: 1,
		}
		err = sqldb.InsertUserRoom(AuserRoom)
		if err != nil {
			log.Println("创建UserRoom记录失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		// 3. UserRoom
		BuserRoom := &models.UserRoom{
			UserID:   user.UserID,
			RoomID:   room.RoomID,
			RoomType: 1,
		}
		err = sqldb.InsertUserRoom(BuserRoom)
		if err != nil {
			log.Println("创建UserRoom记录失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "添加成功",
		})
	}
}
