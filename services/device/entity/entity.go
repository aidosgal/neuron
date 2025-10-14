package entity

import "github.com/aidosgal/neuron/ent"

type (
	Device struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DeviceToken string `json:"token"`
	}

	CreateRequest struct {
		AdminToken string `json:"token"`
		Name       string `json:"name"`
	}

	CreateResponse struct {
		Device Device `json:"device"`
	}
)

func MakeStorageDeviceToEntity(device *ent.Device) *Device {
	return &Device{
		ID: device.ID.String(),
		Name: device.DeviceName,
	}
}
