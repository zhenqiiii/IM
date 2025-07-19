package snowflakeID

import "github.com/bwmarrin/snowflake"

// 节点
var node *snowflake.Node

func GenID() (user_id string) {
	// 设置起始时间和机器ID

	// 先使用默认的起始时间，节点ID设为1
	node, _ = snowflake.NewNode(1)
	// 返回生成ID
	return node.Generate().String()
}
