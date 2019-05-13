package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	client *mongo.Client
}

func NewSession(url string) (*Session, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	return &Session{client}, nil
}

func (s *Session) Close() error {
	return s.client.Disconnect(context.TODO())
}

func (s *Session) DropDatabase(db string) error {
	db := s.client.Database(db)

	return db.DropDatabase(context.TODO())
}
