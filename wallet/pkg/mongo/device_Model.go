package mongo

import (
	root "wallet/pkg"
)

type deviceModel struct {
	Identifier string

	PublicKey  string
	PrivateKey string
}

func newDeviceModel(d *root.Device) (*deviceModel, error) {
	// add code to generate public private key pair
	device := deviceModel{Identifier: d.Identifier, PublicKey: "r1", PrivateKey: "r2"}

	return &device, nil
}
