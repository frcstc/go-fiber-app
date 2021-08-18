package middleware

import (
	"fiber/app/repository/userRepository"
	businessError "fiber/error"
	"fiber/global"
	"fiber/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware jwt 校验中间件
func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 获取token
		authorization := string(ctx.Request().Header.Peek(fiber.HeaderAuthorization))
		authArr := strings.Split(authorization, " ")
		if len(authArr) != 2 || authArr[0] != "Bearer" {
			panic(businessError.New(businessError.TOKEN_NOT_EXISTS))
		}
		tokenString := authArr[1]

		// 解析token
		customClaims, err := utils.ParseToken(tokenString)
		if err != nil {
			panic(businessError.New(businessError.TOKEN_PARSE_ERROR))
		}

		// 获取用户的信息
		user := userRepository.GetUserById(customClaims.UserId)
		if user == nil {
			panic(businessError.New(businessError.USER_NOT_FOUND))
		}

		// 判断密码版本号
		if user.PasswordVersion != customClaims.PasswordVersion {
			panic(businessError.New(businessError.TOKEN_EXPIRED))
		}

		// 放入上下文中
		global.SetAuthUser(user)

		return ctx.Next()
	}
}


