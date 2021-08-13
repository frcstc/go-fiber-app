/**
 * @Author: AF
 * @Date: 2021/8/10 12:03
 */

package logger

import (
	"fiber/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func init() {
	logFilePath := "log/"
	logFileName := "app.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	_ = os.Mkdir(logFilePath, 0755)
	global.LogFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.Out = global.LogFile
	logrus.SetLevel(log.TraceLevel)
	writer, _ := rotatelogs.New(
		logFilePath+"%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  writer,
		log.FatalLevel: writer,
		log.DebugLevel: writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.PanicLevel: writer,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(lfHook)
	global.SLog = logrus
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		global.BLog = global.SLog.WithFields(log.Fields{
			"requestId": reqId,
			"requestIp": c.IP(),
		})
		return c.Next()
	}
}
