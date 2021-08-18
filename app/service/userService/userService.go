package userService

import (
	userQo "fiber/app/entity/qo"
	"fiber/app/repository/userRepository"
	"fiber/model"
	"fiber/utils"
)

func GetUserByMobile(qo *userQo.GetUserByMobileQo) *model.UserModel {
	return userRepository.GetUserByMobile(qo)
}

func GetPasswordEncrypt(password string, hash string) string  {
	return utils.SHA1(password + hash)
}