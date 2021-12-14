package db

import "time"

type User struct {
    ID          int     `gorm:"primary_key"`
    Username    string  `gorm:"uniqueIndex"`
    Password    string
}

type Device struct {
    ID          int     `gorm:"primary_key"`
    Imei        int     `gorm:"uniqueIndex"`
    Status      int
    IsVirtual   bool    `gorm:"is_virtual"`
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

type Fence struct {
    ID                  int
    SchoolID            int     `gorm:"school_id"`
    SouthwestLongitude  float64 `gorm:"southwest_longitude"`
    SouthwestLatitude   float64 `gorm:"southwest_latitude"`
    NortheastLongitude  float64 `gorm:"northeast_longitude"`
    NortheastLatitude   float64 `gorm:"northeast_latitude"`
}

type LateTime struct {
    ID          int
    SchoolID    int         `gorm:"school_id"`
    Time        string      `gorm:"time"`
}

type LateSituation struct {
    ID          int
    StudentID   int         `gorm:"student_id"`
    Situation   string
    Time        time.Time
}


