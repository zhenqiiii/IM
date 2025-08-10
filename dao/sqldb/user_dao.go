package sqldb

import (
	"log"

	"github.com/zhenqiiii/IM-GO/models"
	"golang.org/x/crypto/bcrypt"
)

/*保存UserBasic相关DAO层函数*/

// 根据Email查询UserBasic
func GetUserBasicByEmail(email string) (user *models.UserBasic, err error) {
	result := db.Where("email = ?", email).Limit(1).First(&user)
	if result.Error != nil {
		log.Println("[DB]GetUserBasicByEmail Error：" + result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}

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

// 校验密码
func VerifyPwd(raw string, userid string) (bool, error) {
	// 查询user_basic
	var user models.UserBasic
	result := db.Where("user_id = ?", userid).Limit(1).Find(&user)
	if result.Error != nil {
		log.Println("[DB]VerifyPwd Error:" + result.Error.Error())
		return false, result.Error
	}
	// 比对
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(raw))
	if err != nil { //false
		log.Println(err)
		return false, nil
	}
	// 密码正确
	return true, nil

}

// 修改密码
func UpdatePwd(rawNew string, userid string) error {
	hashedNew, _ := bcrypt.GenerateFromPassword([]byte(rawNew), bcrypt.DefaultCost)
	result := db.Model(&models.UserBasic{}).Where("user_id = ?", userid).Update("password", hashedNew)
	if result.Error != nil {
		log.Println("[DB]UpdatePwd Error:" + result.Error.Error())
		return result.Error
	}
	return nil
}
