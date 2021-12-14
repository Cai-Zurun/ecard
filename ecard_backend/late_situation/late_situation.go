package late_situation

import (
    "ecard_backend/db"
    "ecard_backend/log"
    "fmt"
    "github.com/robfig/cron"
    "strings"
    "time"
)

const (
    Late    = "未到学校"
    Arrived = "已到学校"
)

/*
判断是否迟到的定时任务
 */
func Daemon() {
    var lateTimes []db.LateTime
    err := db.MySQL.Find(&lateTimes).Error
    if err != nil {
        log.Error("database query error:%v", err)
        return
    }

    c := cron.New()
    for _, lateTime := range lateTimes {
        times := strings.Split(lateTime.Time, ":")
        hour, minute, second := times[0], times[1], times[2]
        spec := fmt.Sprintf("%s %s %s * * ?", second, minute, hour)
        err = c.AddFunc(spec, func() {
            /*
            判断学生是否迟到
             */
            // 周六日跳过
            if time.Now().Weekday().String() == "Saturday" || time.Now().Weekday().String() == "Sunday" {
                return
            }

            // 如果学生最新定位不在学校范围内，则为迟到

            var fence db.Fence
            err = db.MySQL.Where("school_id = ?", lateTime.SchoolID).First(&fence).Error
            if err != nil {
                log.Error("database query error:%v", err)
                return
            }

            var students []db.Student
            err = db.MySQL.Where("school_id = ?", lateTime.SchoolID).Find(&students).Error
            if err != nil {
                log.Error("database query error:%v", err)
                return
            }

            for _, student := range students {
                lateSituation := db.LateSituation{
                    StudentID: student.ID,
                    Time: time.Now(),
                }
                var deviceLocations  []db.Location
                err = db.MySQL.Where("imei = ?", student.Imei).Find(&deviceLocations).Order("created_time desc").Error
                if err != nil {
                    log.Error("database query error:%v", err)
                    continue
                }
                if len(deviceLocations) == 0 {
                    lateSituation.Situation = Late
                    err = db.MySQL.Create(&lateSituation).Error
                    if err != nil {
                        log.Error("database error:%v", err)
                    }
                    continue
                }

                newestLocation := deviceLocations[0]
                if isInFence(newestLocation, fence) == true {
                    lateSituation.Situation = Arrived
                } else {
                    lateSituation.Situation = Late
                }
                err = db.MySQL.Create(&lateSituation).Error
                if err != nil {
                    log.Error("database error:%v", err)
                }
            }
        })
        if err != nil {
            log.Error("late-time cron addFunc error:%v", err)
            return
        }
    }
    c.Start()
}

func isInFence(location db.Location, fence db.Fence) bool {
    if location.Longitude < fence.SouthwestLongitude {
        return false
    }
    if location.Longitude > fence.NortheastLongitude {
        return false
    }
    if location.Latitude < fence.SouthwestLatitude {
        return false
    }
    if location.Latitude > fence.NortheastLatitude {
        return false
    }
    return true
}