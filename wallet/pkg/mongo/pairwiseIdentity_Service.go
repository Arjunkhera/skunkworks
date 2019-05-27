package mongo

import (
	"context"
	"log"
	root "wallet/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PairIdentityService handles the interface of root.Device
type PairIdentityService struct {
	collection *mongo.Collection
}

// NewPairIdentityService creates an instance of  PairIdentityService
func NewPairIdentityService(session *Session, config *root.MongoConfig) *PairIdentityService {
	collection := session.client.Database(config.DbName).Collection("PairIdentity")

	return &PairIdentityService{collection}
}

// CreatePairIdentity creates and inserts an ID into mongo db
func (pId *PairIdentityService) CreatePairIdentity(username string, otherPartyName string) error {
	pairID, err := newPairIdentityModel(username, otherPartyName)
	if err != nil {
		return err
	}

	_, error := pId.collection.InsertOne(context.TODO(), pairID)

	return error
}

func (pId *PairIdentityService) GetPairIdentityByUsername(username string) (root.PairIdentity, error) {
	model := pairIdentityModel{}
	filter := bson.D{{"username", username}}

	err := pId.collection.FindOne(context.TODO(), filter).Decode(&model)

	result := root.PairIdentity{
		Identifier: model.Identifier,
		UserName:   model.UserName,
		PublicKey:  model.PublicKey}

	return result, err
}

func (pId *PairIdentityService) GetAllPairIdentities() ([]root.PairIdentity, error) {
	var results []root.PairIdentity

	findOptions := options.Find()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := pId.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var singleId root.PairIdentity
		err := cur.Decode(&singleId)
		if err != nil {
			return nil, err
		}

		results = append(results, singleId)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results, nil
}
