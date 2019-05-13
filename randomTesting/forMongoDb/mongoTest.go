package main

import (
	"context"
	"fmt"
	"log"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	Identifier  string `json:"ID"`
	CommonName  string `json:"CommonName"`
	PhoneNumber string `json:"PhoneNumber"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
	GetRecordByIdentifier(identifier string) (error, Record)
}

func main() {

	// connect to the database
	client := createConnection()

	// check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to the Database")

	collection := client.Database("test").Collection("Records")

	firstRecord := Record{Identifier: "1", CommonName: "ArjunKhera", PhoneNumber: "9999485949"}
	// insert a record into the database
	insertResult, err := collection.InsertOne(context.TODO(), firstRecord)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func createConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func dropDatabase() {
}

func closeConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
