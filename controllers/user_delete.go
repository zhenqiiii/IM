package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 删除好友处理函数
func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收参数
		id_B := c.PostForm("user_id")
		if id_B == "" {
			log.Println("请求携带的user_id参数为空")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "用户user_id参数为空",
			})
			return
		}

		// 删除三条记录：RoomBasic*1 + UserRoom*2
		// 获取二者所在房间RoomID
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		roomid, err := sqldb.GetTwoUsersRoom(userInfo.UserID, id_B)
		if err != nil {
			log.Println("获取房间RoomID失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常:" + err.Error(),
			})
			return
		}
		// delete
		err = sqldb.DeleteRoomBasic(roomid)
		if err != nil {
			log.Println("删除RoomBasic失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常:" + err.Error(),
			})
			return
		}
		err = sqldb.DeleteUserRoom(roomid)
		if err != nil {
			log.Println("删除UserRoom失败：" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统异常:" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "删除成功",
		})
	}
}
