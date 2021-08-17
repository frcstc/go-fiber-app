/**
 * @Author: AF
 * @Date: 2021/8/9 14:53
 */

package router

import (
	"fiber/app/api"
	businessError "fiber/error"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
)

func AppRouter(app *fiber.App) {
	// 路由设置
	app.Post("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx))
	})
	app.Post("/LoginByPassword", auth.PasswordLogin)
	// 404返回
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(resultVo.Fail(businessError.New(businessError.NOT_FOUND), c))
	})
}
