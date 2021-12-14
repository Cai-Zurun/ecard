package handler

import (
	"ecard_gateway/db"
	"ecard_gateway/ecard/conn_map"
	"ecard_gateway/ecard/packet"
	"ecard_gateway/log"
	"ecard_gateway/utils"
	"fmt"
	"time"
)

func handleGps(packet *packet.Packet) []byte {
	location := db.Location{}

	var err error
	imei := conn_map.GetDeviceInfo(packet.Conn).Imei
	location.Imei = imei
	location.CreatedTime = parseTime(packet.PacketBuf[4:11])
	location.Latitude = utils.BytesToDecimal(packet.PacketBuf[11:15]) / 10000000.0
	location.Longitude = utils.BytesToDecimal(packet.PacketBuf[15:19]) / 10000000.0
	location.LocationStr, err = utils.Regeo(location.Latitude, location.Longitude)
	if err != nil {
		log.Error("地址逆解析失败:%v", err)
		return nil
	}
	return nil
}

func parseTime(buf []byte) time.Time {
	year := int(utils.BytesToDecimal(buf[0:1]))
	month := int(utils.BytesToDecimal(buf[1:2]))
	day := int(utils.BytesToDecimal(buf[2:3]))
	hour := int(utils.BytesToDecimal(buf[3:4]))
	minute := int(utils.BytesToDecimal(buf[4:5]))
	second := int(utils.BytesToDecimal(buf[5:6]))

	t, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("20%d-%d-%d %d:%d:%d", year,month,day,hour,minute,second))
	if err != nil {
		log.Error("time parse error:%v", err)
		return time.Now()
	}
	return t
}
