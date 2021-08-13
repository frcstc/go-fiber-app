/**
 * @Author: AF
 * @Date: 2021/8/11 10:09
 */

package repository

import (
	"fiber/global"
	"fiber/model"
)

// GetUserByMobile 手机号和 手机号前缀匹配
func GetUserByMobile(mobile string) model.UserModel {
	var user model.UserModel
	global.DB.Where("mobile = ?", mobile).First(&user)
	return user
}
