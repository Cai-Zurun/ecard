package controller

import (
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/db"
    "github.com/gin-gonic/gin"
)

/*
获取学校信息
 */
func GetSchoolInfo(c *gin.Context) {
    var schools []db.School
    err := db.MySQL.Find(&schools).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    SuccessResponse(c, global.SUCCESS, schools)
    return
}
