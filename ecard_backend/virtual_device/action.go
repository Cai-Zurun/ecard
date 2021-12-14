package virtual_device

import (
    "ecard_backend/log"
    "ecard_backend/utils"
    "encoding/hex"
    "net"
    "strconv"
    "time"
)

func Login(conn net.Conn, imei int) error {
    loginPacket := []byte{0x78, 0x78, 0x0d, 0x01}

    imeiStr := strconv.Itoa(imei)
    imeiByte, err := hex.DecodeString(imeiStr)
    if err != nil {
        log.Error("imei decode string error:%v", err)
        return err
    }
    loginPacket = append(loginPacket, imeiByte...)

    deviceStuff := GetDeviceStuff(imei)
    deviceStuff.SeqNum++
    loginPacket = append(loginPacket, utils.IntToSeqBytes(deviceStuff.SeqNum)...)
    loginPacket = append(loginPacket, utils.CrcBytes(loginPacket[2:])...)
    loginPacket = append(loginPacket, []byte{0x0d, 0x0a}...)

    _, err = conn.Write(loginPacket)
    if err != nil {
        log.Error("conn.Write error:", err.Error())
        return err
    }

    time.Sleep(2*time.Second)
    // 等待服务器返回登录响应包
    //loginResponse := make([]byte, 10)
    //_, err = conn.Read(loginResponse)
    //if err != nil {
    //    log.Error("conn.Read error:", err.Error())
    //    return err
    //}
    //if len(loginResponse) < 4 || !bytes.Equal(loginResponse[:4], []byte{0x78, 0x78, 0x05, 0x01}){
    //    err = errors.New("wrong loginResponse")
    //    log.Error(err.Error())
    //    return err
    //}
    return nil
}


func HeartBeating(conn net.Conn, imei int) {
    deviceStuff := GetDeviceStuff(imei)

    ticker := time.NewTicker(5*time.Minute)
    defer ticker.Stop()

    for {
        <-ticker.C

        heartPacket := []byte{0x78, 0x78, 0x05, 0x23}
        deviceStuff.SeqNum++
        heartPacket = append(heartPacket, utils.IntToSeqBytes(deviceStuff.SeqNum)...)
        heartPacket = append(heartPacket, utils.CrcBytes(heartPacket[2:])...)
        heartPacket = append(heartPacket, []byte{0x0d, 0x0a}...)

        _, err := conn.Write(heartPacket)
        if err != nil {
            log.Error("conn.Write error:", err.Error())
        }
    }
}


func Alarm(conn net.Conn, imei int) {
    deviceStuff := GetDeviceStuff(imei)
    alarmChannel := deviceStuff.AlarmChannel

    for {
        select {
        case alarmType := <-alarmChannel:
            alarmPacket := []byte{0x78, 0x78, 0x06, 0x27}
            alarmPacket = append(alarmPacket, alarmType)
            deviceStuff.SeqNum++
            alarmPacket = append(alarmPacket, utils.IntToSeqBytes(deviceStuff.SeqNum)...)
            alarmPacket = append(alarmPacket, utils.CrcBytes(alarmPacket[2:])...)
            alarmPacket = append(alarmPacket, []byte{0x0d, 0x0a}...)

            _, err := conn.Write(alarmPacket)
            if err != nil {
                log.Error("conn.Write error:", err.Error())
            }
        }
    }
}

func Locate(conn net.Conn) {

}

