package qmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	info := Info{
		Host:     "localhost",
		Port:     "27017",
		Account:  "root",
		Password: "example",
	}
	client, err := NewClient(&info)
	if err != nil {
		t.Log(err)
		return
	}

	collection := client.Database("testing").Collection("numbers")

	res, err := collection.InsertOne(context.TODO(), bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.InsertedID)
}
