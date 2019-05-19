package mongo

import (
	"encoding/json"
	"io/ioutil"
	root "wallet/pkg"
)

// DeviceService handles the interface of root.Device
type DeviceService struct {
	filePath string
}

// NewDeviceService creates an instance of  UserService
func NewDeviceService(path string) *DeviceService {
	return &DeviceService{filePath: path}
}

// CreateDevice stores device information in the form of json
func (devServ *DeviceService) CreateDevice(device *root.Device) error {
	dev, err := newDeviceModel(device)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(dev, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(devServ.filePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
