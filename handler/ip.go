package handler

import (
	"log"
)

var _devices map[string]STDevice

//STDevice ..
type STDevice struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

func init() {
	_devices = make(map[string]STDevice)
}

//SaveDeviceMessage ..
func SaveDeviceMessage(device string, host string) bool {
	if len(device) < 2 {
		return false
	}
	log.Println("device:" + device + " - " + host)

	var d STDevice
	d.Name = device
	d.Host = host

	_devices[device] = d
	return true
}

//GetDevices ..
func GetDevices() string {

	str := ""
	for _, keyB := range _devices {
		str += keyB.Name + " - " + keyB.Host + "; "
	}

	return str
}
