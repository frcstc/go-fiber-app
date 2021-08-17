package utils

import (
	"fiber/config"
	"fiber/global"
	"fiber/resultVo"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

type CustomClaims struct {
	UserId          string `json:"userId"`
	PasswordVersion int8    `json:"passwordVersion"`
	jwt.StandardClaims
}

func CreateToken(userId string, passwordVersion int8) string {
	customClaims := &CustomClaims{
		UserId:          userId,
		PasswordVersion: passwordVersion,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TTL).Unix(),
			Id: utils.UUID(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		global.BLog.Error("jwt token generate err", err)
		panic(resultVo.TOKEN_CREATE_ERROR)
	}
	global.SLog.Infof("user login success:  userId: %s token: %s", userId, tokenString)
	return tokenString
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT_SECRET), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
