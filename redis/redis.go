package redis

import (
	"fiber/config"
	"fiber/global"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func InitRedis() {
	// 默认连接池
	defaultPool := defaultPool()
	global.Redis = defaultPool.Get()
}

func defaultPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			db, _ := strconv.Atoi(config.Config("REDIS_DB"))
			conn, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:6379", config.Config("REDIS_HOST")),
				redis.DialReadTimeout(time.Second),
				redis.DialWriteTimeout(time.Second*2),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialDatabase(db),
				redis.DialPassword(config.Config("REDIS_PASSWORD")),
			)
			if err != nil {
				global.SLog.Errorf("redis conn err %v", err)
				fmt.Println("redis conn err", err)
			}
			return conn, err
		},
		MaxIdle:     50,
		MaxActive:   10000,
		IdleTimeout: 60 * time.Second,
	}
}
