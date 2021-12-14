package controller

import (
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/db"
    "ecard_backend/utils"
    "github.com/gin-gonic/gin"
)

func AddDevice(c *gin.Context) {
    // 获取请求参数
    var addDeviceRequest global.AddDeviceRequest
    err := c.ShouldBindJSON(&addDeviceRequest)
    if err != nil {
       log.Error("Bind Json Error:%s", err.Error())
       FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
       return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&addDeviceRequest)
    if err != nil {
       log.Error("Request Validate Error:%v", err)
       FailureResponse(c, global.BAD_REQUEST, err.Error())
       return
    }

    // 添加设备
    err = db.MySQL.Create(&db.Device{
        Imei: addDeviceRequest.Imei,
        Status: 0,
    }).Error
    if err != nil {
       log.Error("AddDevice Error:%v", err)
       FailureResponse(c, global.BAD_REQUEST, err.Error())
       return
    }

    SuccessResponse(c, global.SUCCESS, "添加设备成功")
    return
}

/*
获取设备总数、在线设备数、离线设备数
 */
func GetDeviceCount(c *gin.Context) {
    var deviceCount, onlineCount, offlineCount int
    err := db.MySQL.Model(&db.Device{}).Count(&deviceCount).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    err = db.MySQL.Model(&db.Device{}).Where("status = ?", 1).Count(&onlineCount).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    err = db.MySQL.Model(&db.Device{}).Where("status = ?", 0).Count(&offlineCount).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    SuccessResponse(c, global.SUCCESS, global.GetDeviceCountResponse{
        DeviceCount: deviceCount,
        OnlineCount: onlineCount,
        OfflineCount: offlineCount,
    })
    return
}
