package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoCN = ConnectorDB()
	// clientOpts = options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(options.Credential{
	// 	Username: "root",
	// 	Password: "password",
	// })
	clientOpts = getEnvOptions()
)

func ConnectorDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOpts)
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

func getEnvOptions() *options.ClientOptions {
	if os.Getenv("MONGO_URI") == "" {
		return options.Client().ApplyURI("mongodb://localhost:32717")
		// return options.Client().ApplyURI("mongodb://localhost:27017")
	}
	return options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(options.Credential{
		Username: "admin",
		Password: "Passw0rd",
	})
	// return options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(options.Credential{
	// 	Username: "root",
	// 	Password: "password",
	// })
}
