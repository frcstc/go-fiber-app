package userRepository

import (
	userQo "fiber/app/entity/qo"
	"fiber/global"
	"fiber/model"
)

func GetUserByMobile(qo *userQo.GetUserByMobileQo) *model.UserModel {
	var user = new(model.UserModel)
	tx := global.DB.Where("mobile = ?", qo.Mobile)
	if qo.MobilePrefix == "" {
		qo.MobilePrefix = "86"
	}
	tx.Where("mobile_prefix", qo.MobilePrefix).First(user)
	return user
}
