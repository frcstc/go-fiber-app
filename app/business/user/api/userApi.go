/**
 * @Author: AF
 * @Date: 2021/8/9 16:43
 */

package api

import (
	"fiber/app/business/user/entity"
	"fiber/app/business/user/service"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
)

type UserApi struct{}

func (user *UserApi) GetUser(c *fiber.Ctx) error {
	var dto entity.RegisterDto

	if err := c.BodyParser(&dto); err != nil {
		panic(resultVo.BAD_REQUEST)
	}
	if dto.Mobile == "" {
		panic(resultVo.BAD_REQUEST)
	}

	result := new(service.UserService).Register(dto)
	return c.JSON(resultVo.Success(result, c))
}
