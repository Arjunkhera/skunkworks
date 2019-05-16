package mongo

import (
	_ "skynet/pkg"
)

type recordModel struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

/*
func newRecordModel() () {
}

type miscellaneousDataModel struct {
	Identifier string `json:"ID"`

	RandomData  string `json:"Key"`
	RandomValue string `json:"Value"`
}
*/
