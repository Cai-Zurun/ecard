package db

import (
    "ecard_backend/conf"
    "github.com/gomodule/redigo/redis"
    "time"
)

var RedisPool *redis.Pool

func init() {
    RedisPool = &redis.Pool{
        Dial: func() (c redis.Conn, err error) {
            return redis.Dial("tcp", conf.REDIS_HOST, redis.DialPassword(conf.REDIS_PASSWORD))
        },
        MaxIdle:     3,
        IdleTimeout: 240 * time.Second,
    }
}