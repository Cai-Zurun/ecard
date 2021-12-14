package db

import "time"

type DeviceInfo struct {
    Imei    int
}

type User struct {
    ID          int     `gorm:"primary_key"`
    Username    string  `gorm:"uniqueIndex"`
    Password    string
}

type Device struct {
    ID      int     `gorm:"primary_key"`
    Imei    int     `gorm:"uniqueIndex"`
    Status  int
}

type Student struct {
    ID          int     `gorm:"primary_key"`
    StudentNo   int     `gorm:"uniqueIndex"`
    Name        string
    Sex         int
    Age         int
    Phone       string  `gorm:"uniqueIndex"`
    Class       int
    Grade       int
    SchoolID    int
    Imei        int
}

type Alarm struct {
    ID          int
    AlarmType   string      `gorm:"alarm_type"`
    Imei        int
    AlarmTime   time.Time   `gorm:"alarm_time"`
}

type School struct {
    ID          int
    SchoolName  string      `gorm:"school_name"`
}

type Location struct {
    ID          int
    Imei        int
    Longitude   float64
    Latitude    float64
    LocationStr string      `gorm:"location_str"`
    CreatedTime time.Time   `gorm:"created_time"`
}




