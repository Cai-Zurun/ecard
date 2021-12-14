package controller

import (
    "ecard_backend/global"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
    c.JSON(200, Response{
        Code: code,
        Msg:  global.GetMsg(code),
        Data: data,
    })
    c.Abort()
}

func FailureResponse(c *gin.Context, code int, data interface{}) {
    c.JSON(400, Response{
        Code: code,
        Msg:  global.GetMsg(code),
        Data: data,
    })
    c.Abort()
}
