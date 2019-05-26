package mongo

import (
	root "skynet/pkg"
)

type RecordService struct {
	collection *mongo.collection
}

func NewRecordService(session *Session, config *root.MongoConfig) *RecordService {
	collection := session.client.Database(config.DbName).Collection("Record")

	return &RecordService{collection}
}

func (recServ *RecordService) CreateRecord(rec *root.Record, username string) error {

}
