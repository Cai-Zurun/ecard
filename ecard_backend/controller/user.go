package controller

import (
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/service"
    "ecard_backend/utils"
    "github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
    // 获取请求参数
    var loginRequest global.LoginRequest
    err := c.ShouldBindJSON(&loginRequest)
    if err != nil {
        log.Error("Bind Json Error:%s", err.Error())
        FailureResponse(c, global.ERROR_BIND_JSON, err.Error())
        return
    }

    // 校验请求参数
    err = utils.GetValidator().Struct(&loginRequest)
    if err != nil {
        log.Error("Request Validate Error:%v", err)
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    // 登录逻辑
    token, err := service.UserLogin(&loginRequest)
    if err != nil {
        FailureResponse(c, global.BAD_REQUEST, err.Error())
        return
    }

    SuccessResponse(c, global.SUCCESS, global.LoginResponse{
        Token: token,
    })
    return
}
