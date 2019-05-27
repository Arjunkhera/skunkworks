package mongo

import (
	"context"
	root "skynet/pkg"

	"go.mongodb.org/mongo-driver/mongo"
)

type PairIdentityService struct {
	pairIdentityCollection *mongo.Collection
}

func NewPairIdentityService(session *Session, config *root.MongoConfig) *PairIdentityService {
	pairIdentityCollection := session.client.Database(config.DbName).Collection("PairIdentity")

	return &PairIdentityService{pairIdentityCollection}
}

func (pId *PairIdentityService) CreatePairIdentity(newUser root.PairIdentity) error {

	_, err := pId.pairIdentityCollection.InsertOne(context.TODO(), newUser)

	return err
}
