package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/utils"
    "github.com/gin-gonic/gin"
)

func SetLateTime(c *gin.Context) {
    // 获取请求参数
    var setLateTimeRequest global.SetLateTimeRequest
    err := c.ShouldBindJSON(&setLateTimeRequest)
    if err != nil {
        log.Error("Bind Json Error:%s", err.Error())
        FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
        return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&setLateTimeRequest)
    if err != nil {
        log.Error("Request Validate Error:%v", err)
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    var lateTime db.LateTime
    err = db.MySQL.Where("school_id = ?", setLateTimeRequest.SchoolID).First(&lateTime).Error
    if err != nil {
        lateTime.SchoolID = setLateTimeRequest.SchoolID
        lateTime.Time = setLateTimeRequest.LateTime

        err = db.MySQL.Table("late_times").Create(&lateTime).Error
        if err != nil {
            log.Error("database insert error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
    } else {
        lateTime.Time = setLateTimeRequest.LateTime
        err = db.MySQL.Table("late_times").Where("school_id = ?", setLateTimeRequest.SchoolID).Updates(&lateTime).Error
        if err != nil {
            log.Error("database insert error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
    }

    SuccessResponse(c, global.SUCCESS, "设置迟到时间成功")
    return
}

func GetLateTime(c *gin.Context) {
    var lateTime db.LateTime
    err := db.MySQL.Where("school_id = ?", 1).First(&lateTime).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    getLateTimeResponse := global.GetLateTimeResponse{
        LateTime: lateTime.Time,
    }

    SuccessResponse(c, global.SUCCESS, getLateTimeResponse)
    return
}

