package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	root "skynet/pkg"
	"skynet/pkg/crypto"
)

type ClaimService struct {
	claimCollection     *mongo.Collection
	claimDefnCollection *mongo.Collection
}

func NewClaimService(session *Session, config *root.MongoConfig) *ClaimService {
	claimCollection := session.client.Database(config.DbName).Collection("Claim")
	claimDefnCollection := session.client.Database(config.DbName).Collection("ClaimDefn")

	return &ClaimService{claimCollection, claimDefnCollection}
}

func (claimServ *ClaimService) GetClaimByUserID(identifier string) ([]root.Claim, error) {

	var claims []root.Claim
	claim := root.Claim{}
	filter := bson.D{{"useridentifier", identifier}}

	cur, err := claimServ.claimCollection.Find(context.TODO(), filter, options.Find())

	for cur.Next(context.TODO()) {
		err := cur.Decode(&claim)
		if err != nil {
			return nil, err
		}

		claims = append(claims, claim)
	}

	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claims, nil
}

func (claimServ *ClaimService) GetClaimDefnByClaimDefnID(identifier string) ([]root.ClaimDefn, error) {

	var claimDefns []root.ClaimDefn 
	claimDefn := root.ClaimDefn{}
	filter := bson.D{{"claimdefnidentifier" identifier}}
	
	cur, err := claimServ.claimDefnCollection.Find(context.TODO(), filter, options.Find())

	for cur.Next(context.TODO()) {
		err := cur.Decode(&claimDefn)
		if err != nil {
			return nil, err
		}

		claimDefns = append(claims, claim) 
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claimDefns, nil 
}
func (claimServ *ClaimService) CreateClaimDefn(attributesToTypes map[string]string) (string, error) {
	var claimDefn root.ClaimDefn
	var c crypto.Crypto
	claimDefn.ClaimDefnIdentifier = c.GenerateRandomASCIIString(20)
	claimDefn.AttributesToType = attributesToTypes
	_, err = claimServ.claimDefnCollection.InsertOne(context.TODO(), claimDefn)

	return claimDefn.ClaimDefnIdentifier, err
}

func (claimServ *ClaimService) CreateClaim(userID string, claimDefnID string) error {

	//TODO check whether both id's exist or not
	claim := root.Claim{userId, claimDefnID}

	_, err := claimServ.claimCollection.InsertOne(context.TODO(), claim)

	return err
}
