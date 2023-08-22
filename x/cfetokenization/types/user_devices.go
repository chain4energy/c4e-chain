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

func (u Device) GetMeasurement(measurementId uint64) (*Measurement, error) {
	for _, measruement := range u.Measurements {
		if measruement.Id == measurementId {
			return measruement, nil
		}
	}
	return nil, fmt.Errorf("measruement not found")
}

func (u Measurement) GetFulfilledActivePowerSum() uint64 {
	fulfilledActivePowerSum := uint64(0)
	for _, fulfiledActivePower := range u.FulfilledActivePower {
		fulfilledActivePowerSum += fulfiledActivePower.Amount
	}
	return fulfilledActivePowerSum
}
