package controller

import (
    "ecard_backend/db"
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/service"
    "ecard_backend/utils"
    "github.com/gin-gonic/gin"
)

/*
添加学生
 */
func AddStudent(c *gin.Context) {
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

    err = service.AddStudentAndDevice(&addStudentRequest, 0, false)
    if err != nil {
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    SuccessResponse(c, global.SUCCESS, "添加学生成功")
    return
}

/*
获取学生总数
 */
func GetStudentCount(c *gin.Context) {
    var studentCount int
    err := db.MySQL.Model(&db.Student{}).Count(&studentCount).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }
    SuccessResponse(c, global.SUCCESS, studentCount)
    return
}

/*
获取全部学生信息
*/
func GetStudentInfo(c *gin.Context) {
    var students []db.Student
    err := db.MySQL.Find(&students).Error
    if err != nil {
        log.Error("database query error:%v", err)
        FailureResponse(c, global.ERROR_MYSQL, err.Error())
        return
    }

    studentInfo := make([]global.GetStudentInfoResponse, len(students))
    for i, student := range students {
        var school db.School
        err = db.MySQL.Find(&school).Where("id = ?", student.SchoolID).Error
        if err != nil {
            log.Error("database query error:%v", err)
            FailureResponse(c, global.ERROR_MYSQL, err.Error())
            return
        }
        studentInfo[i] = global.GetStudentInfoResponse{
            Name: student.Name,
            StudentNo: student.StudentNo,
            Sex: student.Sex,
            Age: student.Age,
            Phone: student.Phone,
            Class: student.Class,
            Grade: student.Grade,
            SchoolName: school.SchoolName,
            Imei: student.Imei,
        }
    }
    SuccessResponse(c, global.SUCCESS, studentInfo)
    return
}