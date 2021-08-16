package api

import (
	"fiber/app/entity"
	"fiber/app/service"
	"fiber/resultVo"
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

	user := userService.CheckUser(loginDto)

	if user == nil {
		panic(resultVo.USER_NOT_FOUND)
	}

	return c.JSON(resultVo.Success(user, c))
}
