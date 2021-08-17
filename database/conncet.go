/**
 * @Author: AF
 * @Date: 2021/8/9 14:23
 */

package database

import (
	"fiber/config"
	"fiber/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
		"10s",
	)

	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //关闭默认事务
		PrepareStmt:            true, // 开启缓存预编译，可以提高后续的调用速度
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             2 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Silent,
		}),
	})
	if err != nil {
		global.SLog.Errorf("DB Connect Error %v", err)
		fmt.Println("DB Connect Error", err)
	}
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	SqlDB, _ := global.DB.DB()
	// 设置连接池中最大的闲置连接数
	SqlDB.SetMaxIdleConns(10)
	// 设置数据库的最大连接数量
	SqlDB.SetMaxOpenConns(100)
	// 这是连接的最大可复用时间
	SqlDB.SetConnMaxLifetime(10 * time.Second)
}
