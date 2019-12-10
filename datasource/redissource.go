package datasource

import (
	"fmt"
	conf2 "github.com/birjemin/gin-structure/conf"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
	"time"
)

var redisPool *redis.Pool

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// redis状态
func StatsRedis() redis.PoolStats {
	return redisPool.Stats()
}

// 关闭redis
func CloseRedis() error {
	if redisPool != nil {
		return redisPool.Close()
	}
	return nil
}

// 初始化redis
func init() {
	redisPool = &redis.Pool{
		MaxIdle:     conf2.Int("redis.maxidle"),
		MaxActive:   conf2.Int("redis.maxactive"),
		IdleTimeout: cast.ToDuration(conf2.Int("redis.idletime")) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", conf2.String("redis.host"), conf2.Int("redis.port")),
				redis.DialDatabase(conf2.Int("redis.db")),
				redis.DialPassword(conf2.String("redis.pass")),
			)
		},
	}
	conn := GetRedis()
	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		panic("redis connect failed.")
	}
}
