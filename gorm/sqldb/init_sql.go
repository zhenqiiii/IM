package sqldb

import (
	"log"

	"github.com/spf13/viper"
	"github.com/zhenqiiii/IM-GO/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接实例
var db *gorm.DB

// SQL数据库初始化
func Init_SQL() {
	// 获取dsn
	dsn := viper.GetString("mysql.dsn")

	// 创建连接
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate
	db.AutoMigrate(
		&models.UserBasic{},
		&models.RoomBasic{},
		&models.MessageBasic{},
		&models.UserRoom{})
	// db.Create(&models.UserBasic{
	// 	UserID:   1234,
	// 	Account:  "zwy",
	// 	Password: "1234",
	// })
}
