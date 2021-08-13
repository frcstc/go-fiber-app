/**
 * @Author: AF
 * @Date: 2021/8/11 10:06
 */

package service

import (
	"fiber/app/business/user/entity"
	"fiber/app/business/user/repository"
	"fiber/global"
	"fiber/model"
	"fiber/resultVo"
)

type UserService struct{}

func (u *UserService) Register(dto entity.RegisterDto) *model.UserModel {
	// 查询用户
	user := repository.GetUserByMobile(dto.Mobile)
	if user != (model.UserModel{}) {
		panic(resultVo.RECORD_EXISTS)
	}
	global.BLog.Info("用户注册 手机号 ", dto.Mobile)
	// 创建用户
	return &user
}
