package mongo

import (
	root "skynet/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type recordModel struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

func newRecordModel(rec *root.Record) {*recordModel, error) {
	record := recordModel{Identifier:rec.Identifier,PublicKey: "random123" ,CommonName:rec.CommonName}

	return record, err 
}


