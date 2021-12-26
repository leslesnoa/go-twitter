package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoCN = ConnectorDB()

	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	// clientOptions = options.Client().ApplyURI("mongodb+srv://root:example@localhost:27017")
)

func ConnectorDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection Success to mongo db")
	return client
}

func CheckingConnection() int {
	if err := MongoCN.Ping(context.TODO(), nil); err != nil {
		return 0
	}
	return 1
}
