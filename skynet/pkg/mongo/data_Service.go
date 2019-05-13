package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type RecordService struct {
	collection *mongo.Collection
}

func NewRecordService(session *Session, dbName string, collectionName string) *RecordService {
	collection := session.client.Database(dbName).Collection(collectionName)

}
