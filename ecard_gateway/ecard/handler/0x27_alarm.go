package handler

import (
    "ecard_gateway/db"
    "ecard_gateway/ecard/conn_map"
    "ecard_gateway/ecard/packet"
    "ecard_gateway/ecard/ws_server"
    "ecard_gateway/log"
    "encoding/json"
    "fmt"
    "time"
)

var alarmMap = map[int]string{
    1:  "SOS报警",
    2:  "低电量报警",
    3:  "出围栏报警",
    4:  "进围栏报警",
}

func handleAlarm(p *packet.Packet) []byte {
    alarmType := alarmMap[int(p.PacketBuf[4])]

    imei := conn_map.GetDeviceInfo(p.Conn).Imei

    alarm := db.Alarm{
        AlarmType: alarmType,
        Imei: imei,
        AlarmTime: time.Now(),
    }

    alarmJson, err := json.Marshal(alarm)
    if err != nil {
        log.Error("json marshal error:%v", err)
    }
    _, err = db.RedisPool.Get().Do("LPUSH", "alarm", alarmJson)
    if err != nil {
       log.Error("redis lpush alarm error:%v", err)
    }

    var student db.Student
    err = db.MySQL.Where("imei = ?", imei).First(&student).Error
    if err != nil {
        log.Error("database query error:%v", err)
        return nil
    }

    _, err = ws_server.WsServer.Push("userID", "alarm", fmt.Sprint(student.Name, ": ", alarmType))
    if err != nil {
        log.Error("websocket server push error:%v", err)
        return nil
    }

    return nil
}
