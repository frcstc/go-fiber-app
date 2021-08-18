/**
 * @Author: AF
 * @Date: 2021/8/9 14:53
 */

package router

import (
	"fiber/app/api/authApi"
	businessError "fiber/error"
	"fiber/middleware"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
)

func AppRouter(app *fiber.App) {
	// auth路由
	app.Post("/loginByPassword", authApi.PasswordLogin)
	// 需要登录鉴权的路由
	apiRoute := app.Group("", middleware.AuthMiddleware())
	apiRoute.Get("/userInfo", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx))
	})
	// 其他
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx))
	})
	// 404返回
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(resultVo.Fail(businessError.New(businessError.NOT_FOUND), c))
	})
	// 这个后面不要写
}
