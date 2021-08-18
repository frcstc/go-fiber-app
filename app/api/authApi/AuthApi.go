package authApi

import (
	"fiber/app/entity"
	userQo "fiber/app/entity/qo"
	"fiber/app/service/userService"
	businessError "fiber/error"
	"fiber/resultVo"
	"fiber/utils"
	"github.com/gofiber/fiber/v2"
)

//PasswordLogin  用户密码登录
func PasswordLogin(c *fiber.Ctx) error {
	var loginDto entity.PasswordLoginDto

	if err := c.BodyParser(&loginDto); err != nil {
		panic(businessError.New(businessError.BAD_REQUEST))
	}

	if loginDto.Password == "" || loginDto.Mobile == "" {
		panic(businessError.New(businessError.PARAMS_ERROR))
	}

	// 用户查询
	qo := &userQo.GetUserByMobileQo{Mobile: loginDto.Mobile, MobilePrefix: loginDto.MobilePrefix}
	user := userService.GetUserByMobile(qo)

	// 判断用户是否存在
	if user == nil {
		panic(businessError.New(businessError.USER_NOT_FOUND))
	}

	// 判断用户是否被锁定
	if user.IsLogout || user.IsLogout {
		panic(businessError.New(businessError.USER_LOCKED))
	}

	// 判断密码是否相等
	passwordStr := userService.GetPasswordEncrypt(loginDto.Password, user.PasswordSalt)
	if passwordStr != user.Password {
		panic(businessError.New(businessError.USER_LOGIN_ERROR))
	}

	// 创建token
	token := utils.CreateToken(user.Id, user.PasswordVersion)
	return c.JSON(resultVo.Success(token, c))
}
