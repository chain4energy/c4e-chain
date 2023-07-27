package types

import "fmt"

func (u UserDevices) GetDevice(deviceAddress string) (*UserDevice, error) {
	for _, device := range u.Devices {
		if device.DeviceAddress == deviceAddress {
			return device, nil
		}
	}
	return nil, fmt.Errorf("device not found")
}
