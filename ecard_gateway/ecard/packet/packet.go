package packet

import (
    "ecard_gateway/utils"
    "encoding/hex"
    "errors"
    "net"
)

type Packet struct {
    PacketBuf []byte // 整个数据包
    //ClientAddr string	// 客户端地址
    Conn          net.Conn // 客户端和服务器的连接
    PacketLen     int      // 包长度
    PacketLenSize int      // 包长度字段的长度
}

/*
根据包长度获取对应的数据包对象
*/
func GetPacket(conn net.Conn) (p *Packet, err error) {
    p = new(Packet)
    p.Conn = conn
    // 读取起始位，判断包长度字段的长度
    p.PacketBuf = make([]byte, 2)
    if err = utils.RecvAll(conn, p.PacketBuf); err != nil {
        return nil, err
    }
    // 包长度字段的长度
    if start := hex.EncodeToString(p.PacketBuf[:2]); start == "7878" {
        p.PacketLenSize = 1
    } else {
        return nil, errors.New("receive unexpected packet")
    }
    // 读取包长度
    packetLenBuf := make([]byte, p.PacketLenSize)
    if err = utils.RecvAll(conn, packetLenBuf); err != nil {
        return nil, err
    }
    // 包长度赋值
    if p.PacketLenSize == 1 {
        p.PacketLen = int(packetLenBuf[0])
    }

    // 将包长度拼接到packetBuf后
    p.PacketBuf = append(p.PacketBuf, packetLenBuf...)
    
    // 读取包长度后面的所有信息
    packetDetailBuf := make([]byte, p.PacketLen)
    if err = utils.RecvAll(conn, packetDetailBuf); err != nil {
        return
    }
    //将包长度字段后的所有信息拼接到packetBuf后
    p.PacketBuf = append(p.PacketBuf, packetDetailBuf...)
    
    // 停止位的2字节
    packetStopBuf := make([]byte, 2)
    if err = utils.RecvAll(conn, packetStopBuf); err != nil {
        return nil, err
    }
    if hex.EncodeToString(packetStopBuf) != "0d0a" {
        return nil, errors.New("receive unexpected packet")
    }
    p.PacketBuf = append(p.PacketBuf, packetStopBuf...)
    return
}