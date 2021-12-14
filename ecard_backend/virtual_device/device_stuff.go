package virtual_device

import "sync"

var deviceMap sync.Map

type DeviceStuff struct {
    SeqNum       int
    AlarmChannel chan byte
}

func AddDevice(imei int) {
    deviceStuff := &DeviceStuff{
        SeqNum:       0,
        AlarmChannel: make(chan byte, 1000),
    }
    deviceMap.Store(imei, deviceStuff)
}

func GetDeviceStuff(imei int) *DeviceStuff {
    deviceStuff, _ := deviceMap.Load(imei)
    return deviceStuff.(*DeviceStuff)
}
