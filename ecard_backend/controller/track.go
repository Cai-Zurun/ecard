package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "github.com/gin-gonic/gin"
)

var (
	locationNum = 5
)

func GetTrackInfo(c *gin.Context) {
    var devices []db.Device
    err := db.MySQL.Find(&devices).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    trackInfo := make([]global.GetTrackInfoResponse, len(devices))
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
            trackInfo[i].StudentName = student.Name
            trackInfo[i].Imei = device.Imei
            continue
        }
        if len(deviceLocations) > 5 {
            deviceLocations = deviceLocations[:5]
        }

        trackInfo[i].StudentName = student.Name
        trackInfo[i].Imei = device.Imei
        trackInfo[i].Track = make([][2]float64, len(deviceLocations))
        for j := 0; j < len(deviceLocations); j++ {
            trackInfo[i].Track[j][0] = deviceLocations[j].Longitude
            trackInfo[i].Track[j][1] = deviceLocations[j].Latitude
            trackInfo[i].TrackStr = append(trackInfo[i].TrackStr, deviceLocations[j].LocationStr)
        }
    }
    SuccessResponse(c, global.SUCCESS, trackInfo)
    return
}