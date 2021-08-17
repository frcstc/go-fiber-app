/**
 * @Author: AF
 * @Date: 2021/8/10 15:02
 */

package middleware

import (
	businessError "fiber/error"
	"fiber/global"
	"fiber/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"os"
	"runtime"
)

func InitMiddleware(app *fiber.App) {
	// 全局ID
	app.Use(requestid.New())
	// 跨域中间件
	app.Use(cors.New())
	// 设置日志中间件
	app.Use(logger.New())
	// 设置全局错误
	app.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(e interface{}) {
			// 如果不是业务异常，就抛出堆栈信息
			if _, ok := e.(*businessError.Err); !ok {
				buf := make([]byte, 1024)
				buf = buf[:runtime.Stack(buf, false)]
				global.SLog.Errorf("runtime error : %v stack: %s", e, buf)
				_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
			}
		},
	}))

}
