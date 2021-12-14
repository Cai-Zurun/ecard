package handler

import (
	"ecard_gateway/db"
	"ecard_gateway/ecard/conn_map"
	"ecard_gateway/ecard/packet"
	"encoding/hex"
	"strconv"
)

func handleLogin(packet *packet.Packet) []byte {
	// 存储登录包中的数据(Imei、Language)到临时的SafeMap中
	device := new(db.DeviceInfo)

	device.Imei, _ = strconv.Atoi(hex.EncodeToString(packet.PacketBuf[4:12]))
	conn_map.AddTcpConn(packet.Conn, device)

	return nil
}
