package mongo

import (
	"skynet/pkg"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recordModel struct {
	Identifier primitive.ObjectId `bson:"_id,omitempty"`
	PublicKey  string             `json:"PubKey"`
	Password   string             `json:"Psw"`
	CommonName string             `json:"CommonName"`
}

func newRecordModel(rec *root.Record) *recordModel {

	return &recordModel{
		PublicKey:  rec.PublicKey,
		Password:   rec.Password,
		CommonName: rec.CommonName}
}

func (rec *recordModel) toRootData() *root.Record {

	return &root.Record{
		Identifier: rec.Identifier.Hex(),
		PublicKey:  rec.PublicKey,
		Password:   rec.Password,
		CommonName: rec.CommonName}
}
