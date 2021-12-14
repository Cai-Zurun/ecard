package service

import (
    "crypto/md5"
    "ecard_backend/global"
    "ecard_backend/log"
    "ecard_backend/midlleware"
    "ecard_backend/db"
    "errors"
    "fmt"
)

const (
    SALT = "fdsafagfdgv43532ju76jM"
)

func UserLogin(loginRequest *global.LoginRequest) (string, error) {
    // 检查用户是否存在
    var user db.User
    err := db.MySQL.Where("username = ?", loginRequest.UserName).Find(&user).Error
    if err != nil {
        log.Error("select user error:%v", err)
        return "", err
    }

    // 检查密码是否正确
    encPassword := fmt.Sprintf("%x", md5.Sum([]byte(loginRequest.Password+SALT)))
    if user.Password != encPassword {
        return "", errors.New("密码错误")
    }

    // 验证通过，返回token
    token, err := midlleware.GenerateToken(loginRequest.UserName)
    if err != nil {
        log.Error("GenerateToken Error:%v", err)
        return "", errors.New("GenerateToken Error:"+err.Error())
    }
    return token, nil
}
