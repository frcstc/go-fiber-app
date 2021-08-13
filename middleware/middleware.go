/**
 * @Author: AF
 * @Date: 2021/8/10 15:02
 */

package middleware

import (
	"fiber/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitMiddleware(app *fiber.App) {
	// 全局ID
	app.Use(requestid.New())
	// 跨域中间件
	app.Use(cors.New())
	// 设置日志中间件
	app.Use(logger.New())
	// 设置全局错误
	app.Use(recover2.New())

}
