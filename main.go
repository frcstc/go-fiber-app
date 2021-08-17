/**
 * @Author: AF
 * @Date: 2021/8/9 9:54
 */

package main

import (
	"fiber/config"
	"fiber/database"
	businessError "fiber/error"
	"fiber/global"
	"fiber/middleware"
	"fiber/redis"
	"fiber/resultVo"
	"fiber/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	defer global.CloseGlobal()
	// 配置文件加载
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	// 启动配置
	initConfig := fiber.Config{
		AppName: config.Config("APP_NAME"),
		// 业务异常返回
		ErrorHandler: func(ctx *fiber.Ctx, e error) error {
			if err, ok := e.(*businessError.Err); ok {
				// 业务异常
				return ctx.JSON(resultVo.Fail(err, ctx))
			} else  {
				// 系统异常
				global.BLog.Errorf("server error %v", e)
				return ctx.JSON(resultVo.Fail(businessError.New(businessError.SERVER_ERROR), ctx))
			}
		},
	}

	// 配置初始化
	app := fiber.New(initConfig)
	// 中间件初始化
	middleware.InitMiddleware(app)
	// 初始化数据库连接
	database.ConnectDB()
	// redis
	redis.InitRedis()
	// 初始化路由
	router.AppRouter(app)
	// 启动服务
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config("PORT"))))
}
