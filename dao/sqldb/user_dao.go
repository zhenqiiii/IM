package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
)

/*保存UserBasic相关DAO层函数*/

// 根据用户id查询UserBasic
func GetUserBasicByID(userid string) (user *models.UserBasic, err error) {
	result := db.Where("user_id = ?", userid).Limit(1).First(&user)
	if result.Error != nil {
		log.Println("查询失败：" + result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}

// 根据account查找UserBasic
func GetUserBasicByAccount(account string) (user *models.UserBasic, err error) {
	result := db.Where("account = ?", account).Limit(1).First(&user)
	if result.Error != nil {
		log.Println("查询失败：" + result.Error.Error())
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

// 根据account查询用户是否存在
func CheckUserBasicExistByAccount(account string) (bool, error) {
	result := db.Where("account = ?", account).Limit(1).Find(&models.UserBasic{})
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

// 插入用户
func InsertUserBasic(user models.UserBasic) error {
	result := db.Create(&user)
	if result.Error != nil {
		log.Println("插入用户数据失败：" + result.Error.Error())
		return result.Error
	}
	return nil
}

// 修改资料
func UpdateProfile(userid string, profile models.EditableProfileParams) error {
	result := db.Model(&models.UserBasic{}).Where("user_id = ?", userid).Update("nickname", profile.Nickname).Update("avatar", profile.Avatar).Update("gender", profile.Gender)
	if result.Error != nil {
		log.Println("[DB]UpdateProfile Error:" + result.Error.Error())
		return result.Error
	}
	return nil
}
