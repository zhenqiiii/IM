package sql

import "github.com/zhenqiiii/IM-GO/models"

/*保存UserBasic相关DAO层函数*/

// 根据account查找UserBasic
func GetUserBasicByAccount(account string) (user *models.UserBasic, err error) {
	result := db.Where("account = ?", account).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
