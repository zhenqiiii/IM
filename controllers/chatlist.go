package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/IM-GO/cont"
	"github.com/zhenqiiii/IM-GO/dao/sqldb"
	"github.com/zhenqiiii/IM-GO/pkg/jwt"
)

// 拉取聊天记录列表函数
// 接收参数：RoomID
func ChatList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取房间号：query参数
		roomID := c.Query("room_id")
		if roomID == "" {
			log.Println("房间号参数为空")
			c.JSON(http.StatusOK, gin.H{
				"code": cont.MISSING_PARAMS,
				"msg":  "房间号不能为空",
			})
			return
		}

		// 判断用户是否属于该房间，如果不属于，后续步骤不执行
		userInfo := c.MustGet("user_claims").(*jwt.UserClaims)
		belong, err := sqldb.GetUserRoomByID(userInfo.UserID, roomID)
		if err != nil {
			log.Println("查询用户房间关系失败" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}
		if !belong {
			c.JSON(http.StatusOK, gin.H{
				"code": cont.WRONG_PARAMS,
				"msg":  "未加入该房间，无法访问",
			})
			return
		}
		// 通过，返回聊天记录
		// 消息列表分页返回，所以需要获取page_index和page_size参数
		pageIndex, _ := strconv.ParseInt(c.Query("page_index"), 10, 0)
		pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 0)
		offset := (pageIndex - 1) * pageSize

		// 获取一页聊天记录，根据pageSize&offset确定长度及起始位置
		messageList, err := sqldb.GetMessageListByRoomID(roomID, int(pageSize), int(offset))
		if err != nil {
			log.Println("消息表访问失败：" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": cont.INTERNAL_ERROR,
				"msg":  "系统错误：" + err.Error(),
			})
			return
		}
		// 获取成功，返回数据
		c.JSON(http.StatusOK, gin.H{
			"code": cont.SUCCESS,
			"msg":  "内容获取成功",
			"data": gin.H{
				"list": messageList,
			},
		})
	}
}
