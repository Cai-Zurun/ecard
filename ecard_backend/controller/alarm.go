package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "sort"
)

var alarmMap = map[int]string{
    1:  "SOS报警",
    2:  "低电量报警",
    3:  "出围栏报警",
    4:  "进围栏报警",
}

func GetAlarmCount(c *gin.Context) {
    var alarmCount int
    err := db.MySQL.Model(&db.Alarm{}).Count(&alarmCount).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    alarms, err := db.RedisPool.Get().Do("LRANGE", "alarm", 0, -1)
    if err != nil {
        log.Error("redis query error:%v", err)
        FailureResponse(c, global.ERROR_REDIS, err.Error())
        return
    }
    alarmCount += len(alarms.([]interface{}))

    SuccessResponse(c, global.SUCCESS, alarmCount)
    return
}

func GetAlarmInfo(c *gin.Context) {
    var alarms []db.Alarm
    err := db.MySQL.Find(&alarms).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    alarmInfoResponse := make([]global.GetAlarmInfoResponse, len(alarms))
    for i, alarm := range alarms {
        var student db.Student
        err := db.MySQL.Where("imei = ?", alarm.Imei).First(&student).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }

        var school db.School
        err = db.MySQL.Where("ID = ?", student.SchoolID).First(&school).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }

        alarmInfoResponse[i].AlarmTime = alarm.AlarmTime.String()
        alarmInfoResponse[i].AlarmType = alarm.AlarmType
        alarmInfoResponse[i].StudentName = student.Name
        alarmInfoResponse[i].Grade = student.Grade
        alarmInfoResponse[i].Class = student.Class
        alarmInfoResponse[i].SchoolName = school.SchoolName
        alarmInfoResponse[i].Phone = student.Phone
    }

    alarmJsons, err := db.RedisPool.Get().Do("LRANGE", "alarm", 0, -1)
    if err != nil {
        log.Error("redis query error:%v", err)
        FailureResponse(c, global.ERROR_REDIS, err.Error())
        return
    }
    for _, alarmJson := range alarmJsons.([]interface{}) {
        var alarm db.Alarm
        err = json.Unmarshal(alarmJson.([]byte), &alarm)
        if err != nil {
            log.Error("json unmarshal error %v", err)
            FailureResponse(c, global.ERROR, err.Error())
            return
        }

        var student db.Student
        err := db.MySQL.Where("imei = ?", alarm.Imei).First(&student).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }

        var school db.School
        err = db.MySQL.Where("ID = ?", student.SchoolID).First(&school).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }

        alarmInfo := global.GetAlarmInfoResponse{
            AlarmTime: alarm.AlarmTime.String(),
            AlarmType: alarm.AlarmType,
            StudentName: student.Name,
            SchoolName: school.SchoolName,
            Grade: student.Grade,
            Class: student.Class,
            Phone: student.Phone,
        }
        alarmInfoResponse = append(alarmInfoResponse, alarmInfo)
    }

    SuccessResponse(c, global.SUCCESS, alarmInfoResponse)
    return
}

type PairList []global.GetAlarmRankResponse

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count }

func GetAlarmRank(c *gin.Context) {
    var alarmRankResponse PairList
    for _, alarm := range alarmMap {
        var alarmCount int
        err := db.MySQL.Model(&db.Alarm{}).Where("alarm_type = ?", alarm).Count(&alarmCount).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
        alarmRankResponse = append(alarmRankResponse, global.GetAlarmRankResponse{
            Alarm: alarm,
            Count: alarmCount,
        })
    }
    sort.Sort(sort.Reverse(alarmRankResponse))

    SuccessResponse(c, global.SUCCESS, alarmRankResponse)
    return
}