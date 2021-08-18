package global

import (
	"fiber/model"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

// Redis 默认redis连接池
var Redis redis.Conn

// DB 数据库
var DB *gorm.DB

// SLog 系统日志
var SLog *log.Logger

// BLog 业务上下文日志
var BLog *log.Entry

var LogFile *os.File

type AuthUserPayload struct {
	UserId string
}
var AuthUser *AuthUserPayload

func SetAuthUser(user *model.UserModel)  {
	AuthUser = &AuthUserPayload{UserId: user.Id}
}