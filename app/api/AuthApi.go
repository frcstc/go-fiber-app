package auth

import (
	"fiber/app/entity"
	userQo "fiber/app/entity/qo"
	"fiber/app/service"
	"fiber/resultVo"
	"fiber/utils"
	"github.com/gofiber/fiber/v2"
)

//PasswordLogin  用户密码登录
func PasswordLogin(c *fiber.Ctx) error {
	var loginDto entity.PasswordLoginDto

	if err := c.BodyParser(&loginDto); err != nil {
		panic(resultVo.BAD_REQUEST)
	}

	if loginDto.Password == "" || loginDto.Mobile == "" {
		panic(resultVo.PARAMS_ERROR)
	}

	// 用户查询
	qo := &userQo.GetUserByMobileQo{Mobile: loginDto.Mobile, MobilePrefix: loginDto.MobilePrefix}
	user := userService.GetUserByMobile(qo)

	// 判断用户是否存在
	if user == nil {
		panic(resultVo.USER_NOT_FOUND)
	}

	// 判断用户是否被锁定
	if user.IsLogout || user.IsLogout {
		panic(resultVo.USER_LOCKED)
	}

	// 判断密码是否相等
	passwordStr := userService.GetPasswordEncrypt(loginDto.Password, user.PasswordSalt)
	if passwordStr != user.Password {
		panic(resultVo.USER_LOGIN_ERROR)
	}

	// 创建token
	token := utils.CreateToken(user.Id, user.PasswordVersion)
	return c.JSON(resultVo.Success(token, c))
}
