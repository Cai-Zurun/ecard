package db

import (
    "ecard_backend/conf"
    "ecard_backend/log"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "os"
)

var MySQL *gorm.DB

func Init() {
    var err error
    MySQL, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8",
        conf.DB_USERNAME,
        conf.DB_PASSWORD,
        conf.DB_HOST,
        conf.DB_NAME,
    ))
    if err != nil {
        log.Error("Failed to connect to MySQL database: %v", err)
        os.Exit(1)
    }

    MySQL.AutoMigrate(
        &User{},
        &Device{},
        &Student{},
        &Alarm{},
        &School{},
        &Location{},
        &Fence{},
        &LateTime{},
        &LateSituation{},
    )
}
