package sql

import (
	"os"

	"github.com/zhenqiiii/IM-GO/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接实例
var db *gorm.DB

// 数据库初始化
func Init_DB() (err error) {
	// 获取dsn
	dsn := os.Getenv("SQL_DSN")

	// 创建连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// AutoMigrate
	db.AutoMigrate(&models.UserBasic{})

	// db.Create(&models.UserBasic{
	// 	Name:     "test",
	// 	PassWord: "123",
	// })

	return nil
}
