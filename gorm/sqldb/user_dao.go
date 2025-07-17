package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
)

/*保存UserBasic相关DAO层函数*/

// 根据account查找UserBasic
func GetUserBasicByAccount(account string) (user *models.UserBasic, err error) {
	result := db.Where("account = ?", account).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// 根据email查询用户是否存在
func CheckUserBasicExistByEmail(email string) (bool, error) {
	result := db.Where("email = ?", email).Find(&models.UserBasic{})
	if result.Error != nil {
		log.Println("查询失败：" + result.Error.Error())
		return false, result.Error
	}
	// exists
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}
