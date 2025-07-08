package main

import (
	"github.com/joho/godotenv"
	"github.com/zhenqiiii/IM-GO/gorm/sql"
	"github.com/zhenqiiii/IM-GO/router"
)

func main() {
	// 加载环境变量
	err := godotenv.Load("./config/.env")
	if err != nil {
		panic(err)
	}

	// 初始化sql
	err = sql.Init_DB()
	if err != nil {
		panic(err)

	}

	// 创建
	r := router.SetupRouter()

	// Run
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
