package conn_map

import (
	"ecard_gateway/db"
	"net"
	"sync"
)
//TODO::改成redis，以防止服务器意外中断，导致连接数据丢失
/*
之所以要此模块，是因为除了登录包，其他数据包都不包含Imei，而GPS等信息需携带Imei才能存入数据库。
所以接收到不是登录包的数据包时，可以根据Packet查到Imei
 */
var safeMap sync.Map


/*
TCP连接与设备信息唯一性绑定
 */
func AddTcpConn(conn net.Conn, device *db.DeviceInfo)  {
	safeMap.Store(conn, device)
}


/*
获取本次TCP长连接生命周期内的设备信息
*/
func GetDeviceInfo(conn net.Conn) *db.DeviceInfo {
	val, ok := safeMap.Load(conn)
	if ok {
		return val.(*db.DeviceInfo)
	}
	return nil
}


/*
根据设备imei获取设备的连接
 */
func GetDeviceConn(imei int) (conn net.Conn) {
	safeMap.Range(func(key, value interface{}) bool {
		if value.(*db.DeviceInfo).Imei == imei {
			conn = key.(net.Conn)
			return false
		}
		return true
	})
	return
}




