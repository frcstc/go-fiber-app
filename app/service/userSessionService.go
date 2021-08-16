package userService

import (
	"fiber/app/entity"
	"fiber/global"
	"fiber/model"
	"fiber/utils"
)

func CheckUser(dto entity.PasswordLoginDto) *model.UserModel {
	var user = new(model.UserModel)
	global.DB.Where("mobile = ?", dto.Mobile).First(user)
	return user
}

func GetPasswordEncrypt(password string, hash string) string  {
	return utils.SHA1(password + hash)
}