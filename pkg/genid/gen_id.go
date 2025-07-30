package genid

import (
	"math/rand/v2"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

// 节点
var node *snowflake.Node

// 使用snowflake算法生成用户ID
func GenID() (user_id string) {
	// 设置起始时间和机器ID

	// 先使用默认的起始时间，节点ID设为1
	node, _ = snowflake.NewNode(1)
	// 返回生成ID
	return node.Generate().String()
}

// 生成RoomID：随机
func GenRoomID() string {
	// 固定seed
	code := strconv.FormatInt(rand.Int64N(1000000)+1000000, 10)
	return code
}
