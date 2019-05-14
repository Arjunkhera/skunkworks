package mongo_test

import (
	"log"
	"testing"

	"skynet/pkg"
	"skynet/pkg/mongo"
)

const (
	monogUrl             = "localhost:27017"
	dbName               = "test_db"
	recordCollectionName = "record"
)

func Test_RecordService(t *testing.T) {
	t.Run("Create Record", createAndInsertRecord)
}

func createAndInsertRecord(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongodb %s", err)
	}

	defer func() {
		//session.DropDatabase(dbName)
		session.Close()
	}()

	recServ := NewRecordService(&session, dbName, recordCollectionName)

	testData := root.Record{
		Identifier: "123",
		PublicKey:  "123455",
		Password:   "hello",
		CommonName: "arjun"}

	err = recServ.CreateRecord(&testData)
	if err != nil {
		t.Errorf("Unable to create record %s", err)
	}

	result, err := recServ.GetRecordByIdentifier("123")
	if err != nil {
		t.Errorf("Unable to fetch record %s", err)
	}

	if *result.Password != testData.Password {
		t.Errorf("Incorrect document. Mismatch in data")
	}
}
