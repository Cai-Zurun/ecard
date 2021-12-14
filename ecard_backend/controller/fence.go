package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "strconv"
    "strings"
)


func SetFence(c *gin.Context) {
    // 获取请求参数
    var setFenceRequest global.SetFenceRequest
    err := c.ShouldBindJSON(&setFenceRequest)
    if err != nil {
        log.Error("Bind Json Error:%s", err.Error())
        FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
        return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&setFenceRequest)
    if err != nil {
        log.Error("Request Validate Error:%v", err)
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    southwestLongitude, southwestLatitude, northeastLongitude, northeastLatitude := parseFence(setFenceRequest.Fence)
    newFence := db.Fence{
        SchoolID: setFenceRequest.SchoolID,
        SouthwestLongitude: southwestLongitude,
        SouthwestLatitude: southwestLatitude,
        NortheastLongitude: northeastLongitude,
        NortheastLatitude: northeastLatitude,
    }

    var fence db.Fence
    err = db.MySQL.Where("school_id = ?", setFenceRequest.SchoolID).First(&fence).Error
    if err != nil {
        err = db.MySQL.Table("fences").Create(&newFence).Error
        if err != nil {
            log.Error("database insert error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
    } else {
        err = db.MySQL.Table("fences").Where("school_id = ?", setFenceRequest.SchoolID).Updates(&newFence).Error
        if err != nil {
            log.Error("database insert error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
    }

    SuccessResponse(c, global.SUCCESS, "设置学校范围成功")
    return
}

func GetFence(c *gin.Context) {
    var fence db.Fence
    err := db.MySQL.Where("school_id = ?", 1).First(&fence).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    getFenceResponse := global.GetFenceResponse{
        Fence: fmt.Sprintf("[[%f, %f],[%f, %f]]", fence.SouthwestLongitude, fence.SouthwestLatitude, fence.NortheastLongitude, fence.NortheastLatitude),
    }

    SuccessResponse(c, global.SUCCESS, getFenceResponse)
    return
}

func parseFence(fence string) (southwestLongitude float64, southwestLatitude float64, northeastLongitude float64, northeastLatitude float64) {
    splits := strings.Split(fence, ",")
    for i := 0; i < len(splits); i++ {
        splits[i] = strings.ReplaceAll(splits[i], "[", "")
        splits[i] = strings.ReplaceAll(splits[i], "]", "")
        splits[i] = strings.ReplaceAll(splits[i], " ", "")
    }
    var err error
    southwestLongitude, err = strconv.ParseFloat(splits[0], 64)
    southwestLatitude, err = strconv.ParseFloat(splits[1], 64)
    northeastLongitude, err = strconv.ParseFloat(splits[2], 64)
    northeastLatitude, err = strconv.ParseFloat(splits[3], 64)
    if err != nil {
        log.Error("strconv parse float error:%v", err)
    }
    return
}