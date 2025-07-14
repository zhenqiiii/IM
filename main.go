package main

import (
	"log"

	"github.com/zhenqiiii/IM-GO/gorm/sql"
	"github.com/zhenqiiii/IM-GO/router"
	"github.com/zhenqiiii/IM-GO/utils"
)

func main() {
	// 加载环境变量
	// err := godotenv.Load("./config/.env")
	// if err != nil {
	// 	panic(err)
	// }
	utils.InitConfig()
	// 初始化sql
	sql.Init_SQL()

	// 创建Router
	r := router.SetupRouter()

	// wbskt.GinWebSocketServer()

	// Run
	err := r.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
