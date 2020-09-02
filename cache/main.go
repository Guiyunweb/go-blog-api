package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

// RedisClient Redis缓存客户端单例
var RedisClient redis.Conn

func Redis() {

	db, err := redis.Dial("tcp", os.Getenv("REDIS_ADDR"))

	if err != nil {
		fmt.Println("连接Redis不成功", err)
	}
	RedisClient = db
}
