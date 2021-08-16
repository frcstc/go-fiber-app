package userService

import (
	"fiber/app/entity"
	"fiber/global"
	"fiber/model"
)

func CheckUser(dto entity.PasswordLoginDto) *model.UserModel {
	var user = new(model.UserModel)
	global.DB.Where("mobile = ?", dto.Mobile).First(user)
	return user
}