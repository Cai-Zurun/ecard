package handler

import (
    "ecard_gateway/ecard/packet"
    "ecard_gateway/log"
    "ecard_gateway/utils"
    "encoding/hex"
)

func AnalyzePacket(p *packet.Packet) []byte {
    var protocol string
    if p.PacketLenSize == 1 {
        protocol = hex.EncodeToString(p.PacketBuf[3:4])
    }

    // 检查错误校验码
    calculateCrc := hex.EncodeToString(utils.CrcBytes(p.PacketBuf[2 : len(p.PacketBuf)-4])) //包长度到信息序列号
    realCrc := hex.EncodeToString(p.PacketBuf[len(p.PacketBuf)-4 : len(p.PacketBuf)-2])
    // 校验码错误则抛弃此包
    if calculateCrc != realCrc {
        log.Error("CRC:校验码错误")
        return nil
    }

    // 更新对应imei的序列号
    //deviceInfo := conn_map.GetDeviceInfo(p.Conn)
    //if deviceInfo != nil {
    //    if deviceInfo.Imei != "" {
    //        global.SequenceNumberMap[deviceInfo.Imei] = p.PacketBuf[len(p.PacketBuf)-6 : len(p.PacketBuf)-4]
    //    }
    //}

    switch protocol {
    // 登录包
    case "01":
        return handleLogin(p)
    // 心跳包
    case "23":
        return handleHeart(p)
    // GPS定位包
    case "26":
        return handleGps(p)
    // 报警包
    case "27":
        return handleAlarm(p)
    }
    return nil
}

