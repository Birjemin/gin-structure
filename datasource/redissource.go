package datasource

import (
	"fmt"
	conf "github.com/birjemin/gin-structure/utils/config"
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
		MaxIdle:     conf.Int("redis.maxidle"),
		MaxActive:   conf.Int("redis.maxactive"),
		IdleTimeout: cast.ToDuration(conf.Int("redis.idletime")) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", conf.String("redis.host"), conf.Int("redis.port")),
				redis.DialDatabase(conf.Int("redis.db")),
				redis.DialPassword(conf.String("redis.pass")),
			)
		},
	}
	conn := GetRedis()
	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		panic("redis connect failed.")
	}
}
