package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "github.com/gin-gonic/gin"
)

func GetLateSituation(c *gin.Context) {
    var lateSituations []db.LateSituation
    err := db.MySQL.Order("time desc").Find(&lateSituations).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    getLateSituationResponse := make([]global.GetLateSituationResponse, len(lateSituations))
    for i := 0; i < len(getLateSituationResponse); i++ {
        getLateSituationResponse[i].Situation = lateSituations[i].Situation
        getLateSituationResponse[i].Time = lateSituations[i].Time.String()

        var student db.Student
        err = db.MySQL.Where("id = ?", lateSituations[i].StudentID).First(&student).Error
        if err != nil {
            log.Error("database query error:%v", err)
            continue
        }

        getLateSituationResponse[i].StudentNo = student.StudentNo
        getLateSituationResponse[i].StudentName = student.Name
        getLateSituationResponse[i].Class = student.Class
        getLateSituationResponse[i].Grade = student.Grade
        getLateSituationResponse[i].Phone = student.Phone

        var school db.School
        err = db.MySQL.Where("id = ?", student.SchoolID).First(&school).Error
        if err != nil {
            log.Error("database query error:%v", err)
            continue
        }

        getLateSituationResponse[i].SchoolName = school.SchoolName
    }

    SuccessResponse(c, global.SUCCESS, getLateSituationResponse)
    return
}

