package virtual_device

import (
    "crypto/tls"
    "ecard_backend/conf"
    "ecard_backend/db"
    "ecard_backend/log"
)

func Daemon() {
    // 从数据库获取虚拟设备的数据
    var virtualDevice []db.Device
    err := db.MySQL.Where("is_virtual = ?", true).Find(&virtualDevice).Error
    if err != nil {
        log.Error("database query error:%v", err)
        return
    }

    // 通过协程模拟设备向网关发起连接
    for i := 0; i < len(virtualDevice); i++ {
        log.Info("run ", virtualDevice[i].Imei)
        go RunVirtualDevice(virtualDevice[i].Imei)
    }
}

func RunVirtualDevice(imei int) {
    AddDevice(imei)

    tlsConf := &tls.Config{
        InsecureSkipVerify: true,
    }
    conn, err := tls.Dial("tcp", conf.GATEWAY_ADDR, tlsConf)
    if err != nil {
        log.Error("connect gateway error:%v", err)
        return
    }

    err = Login(conn, imei)
    if err != nil {
        log.Error("send login packet error:%v", err)
        return
    }

    go HeartBeating(conn, imei)
    go Alarm(conn, imei)
}