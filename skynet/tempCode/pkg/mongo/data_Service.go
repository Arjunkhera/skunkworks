package mongo

import (
	"context"
	"skynet/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordService struct {
	collection *mongo.Collection
}

func NewRecordService(session *Session, dbName string, collectionName string) *RecordService {
	collection := session.client.Database(dbName).Collection(collectionName)

	return &RecordService{collection}
}

func (recServ *RecordService) CreateRecord(rec *root.Record) error {
	_, error := recServ.collection.InsertOne(context.TODO(), rec)

	return error
}

func (recServ *RecordService) GetRecordByIdentifier(identifier string) (*root.Record, error) {
	var result root.Record
	filter := bson.D{{"identifier", identifier}}

	err := recServ.collection.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}
