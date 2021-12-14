package db

import (
    "ecard_gateway/conf"
    "ecard_gateway/log"
    "ecard_gateway/utils"
    "encoding/json"
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
    go syncAlarm()
}

func syncAlarm() {
    utils.RecoverAndLog()
    ticker := time.NewTicker(5 * time.Minute)
    defer ticker.Stop()

    for {
        <-ticker.C
        for {
            alarmJson, err := RedisPool.Get().Do("RPOP", "alarm")
            if err != nil {
                log.Error("redis rpop alarm error:%v", err)
                break
            }
            if alarmJson == nil {
                break
            }

            var alarm Alarm
            err = json.Unmarshal(alarmJson.([]byte), &alarm)
            if err != nil {
                log.Error("json unmarshal error %v", err)
            }
            err = MySQL.Create(&alarm).Error
            if err != nil {
                log.Error("database insert error:%v", err)
                break
            }
        }
    }
}







