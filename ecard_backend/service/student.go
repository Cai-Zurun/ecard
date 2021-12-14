package service

import (
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/db"
    "strconv"
)

func AddStudentAndDevice(addStudentRequest *global.AddStudentRequest, status int, isVirtual bool) error {
    studentNo, err :=  strconv.Atoi(addStudentRequest.StudentNo)
    age, err := strconv.Atoi(addStudentRequest.Age)
    class, err := strconv.Atoi(addStudentRequest.Class)
    grade, err := strconv.Atoi(addStudentRequest.Grade)
    imei, err := strconv.Atoi(addStudentRequest.Imei)
    // 添加学生
    err = db.MySQL.Create(&db.Student{
        StudentNo: studentNo,
        Name: addStudentRequest.Name,
        Sex: addStudentRequest.Sex,
        Age: age,
        Phone: addStudentRequest.Phone,
        Class: class,
        Grade: grade,
        SchoolID: addStudentRequest.SchoolID,
        Imei: imei,
    }).Error
    if err != nil {
        log.Error("AddStudent Error:%v", err)
        return err
    }

    // 添加设备
    err = db.MySQL.Create(&db.Device{
        Imei: imei,
        Status: status,
        IsVirtual: isVirtual,
    }).Error
    if err != nil {
        log.Error("AddDevice Error:%v", err)
        return err
    }

    return nil
}
