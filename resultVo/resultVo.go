/**
 * @Author: AF
 * @Date: 2021/8/9 16:48
 */

package resultVo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

type ResultVo struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	TimeStamp int         `json:"timeStamp"`
	RequestId interface{} `json:"requestId"`
	Data      interface{} `json:"data"`
}

func Success(data interface{}, c *fiber.Ctx) ResultVo {
	return ResultVo{
		Code:      SUCCESS,
		Msg:       GetMsg(SUCCESS),
		TimeStamp: time.Now().Nanosecond(),
		Data:      data,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
	}
}

func Fail(code int, c *fiber.Ctx) ResultVo {
	return ResultVo{
		Code:      code,
		Msg:       GetMsg(code),
		TimeStamp: time.Now().Nanosecond(),
		Data:      nil,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
	}
}

func FailCustom(code int, msg string, c *fiber.Ctx) ResultVo {
	return ResultVo{
		Code:      code,
		Msg:       msg,
		TimeStamp: time.Now().Nanosecond(),
		Data:      nil,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
	}
}
