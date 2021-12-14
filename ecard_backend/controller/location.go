package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "github.com/gin-gonic/gin"
)

func GetLocationInfo(c *gin.Context) {
    var devices []db.Device
    err := db.MySQL.Find(&devices).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    locationInfo := make([]global.GetLocationInfoResponse, len(devices))
    for i, device := range devices {
        var student db.Student
        err := db.MySQL.Where("imei = ?", device.Imei).Find(&student).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }

        var deviceLocations  []db.Location
        err = db.MySQL.Where("imei = ?", device.Imei).Order("created_time desc").Find(&deviceLocations).Error
        if err != nil {
            log.Error("database query error:%v", err)
            continue
        }
        if len(deviceLocations) == 0 {
            locationInfo[i].StudentName = student.Name
            locationInfo[i].Imei = device.Imei
            continue
        }

        locationInfo[i].StudentName = student.Name
        locationInfo[i].Imei = device.Imei
        locationInfo[i].Latitude = deviceLocations[0].Latitude
        locationInfo[i].Longitude = deviceLocations[0].Longitude
        locationInfo[i].LocationStr = deviceLocations[0].LocationStr
    }
    SuccessResponse(c, global.SUCCESS, locationInfo)
    return
}