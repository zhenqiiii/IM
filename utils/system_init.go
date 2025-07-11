package utils

import (
	"log"

	"github.com/spf13/viper"
)

// 使用viper初始化config(godotenv也行)
func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")

	// 读入config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// log.Println("sql_dsn:", viper.GetString("dsn"))

	log.Println("Init config successfully")
}
