package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/service"
    "ecard_backend/utils"
    "ecard_backend/virtual_device"
    "github.com/gin-gonic/gin"
    "strconv"
)

func AddVirtualDevice(c *gin.Context) {
    // 获取请求参数
    var addStudentRequest global.AddStudentRequest
    err := c.ShouldBindJSON(&addStudentRequest)
    if err != nil {
        log.Error("Bind Json Error:%s", err.Error())
        FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
        return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&addStudentRequest)
    if err != nil {
        log.Error("Request Validate Error:%v", err)
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    err = service.AddStudentAndDevice(&addStudentRequest, 1,true)
    if err != nil {
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    imei, _ := strconv.Atoi(addStudentRequest.Imei)
    virtual_device.RunVirtualDevice(imei)

    SuccessResponse(c, global.SUCCESS, "添加虚拟设备成功")
    return
}

func Alarm(c *gin.Context) {
    // 获取请求参数
    var alarmRequest global.AlarmRequest
    err := c.ShouldBindJSON(&alarmRequest)
    if err != nil {
        log.Error("Bind Json Error:%s", err.Error())
        FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
        return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&alarmRequest)
    if err != nil {
        log.Error("Request Validate Error:%v", err)
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    deviceStuff := virtual_device.GetDeviceStuff(alarmRequest.Imei)
    deviceStuff.AlarmChannel <- byte(alarmRequest.AlarmType)

    SuccessResponse(c, global.SUCCESS, "报警请求成功")
}

func GetVirtualDevice(c *gin.Context) {
    var virtualDevices []db.Device
    err := db.MySQL.Where("is_virtual = ?", 1).Find(&virtualDevices).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    virtualDeviceResponses := make([]global.GetVirtualDeviceResponse, len(virtualDevices))
    for i := 0; i < len(virtualDeviceResponses); i++ {
        imei := virtualDevices[i].Imei
        virtualDeviceResponses[i].Imei = imei
        virtualDeviceResponses[i].Status = virtualDevices[i].Status

        var student db.Student
        err := db.MySQL.Where("imei = ?", imei).First(&student).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
        virtualDeviceResponses[i].StudentName = student.Name
    }

    SuccessResponse(c, global.SUCCESS, virtualDeviceResponses)
}