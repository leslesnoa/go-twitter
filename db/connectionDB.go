package db

import (
	"context"
	"os"

	"github.com/leslesnoa/go-twitter/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoCN = ConnectorDB()
)

func ConnectorDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), getEnvOptions())
	// client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		logger.Error("Error bad mongoDB connection.", err)
		return client
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logger.Error("Error bad mongoDB connection.", err)
		return client
	}
	logger.Info("Connection success to mongoDB.")
	return client
}

func CheckingConnection() error {
	if err := MongoCN.Ping(context.TODO(), nil); err != nil {
		logger.Error("Error bad mongoDB connection.", err)
		return err
	}
	return nil
}

func getEnvOptions() *options.ClientOptions {
	// username := os.Getenv("MONGO_USERNAME")
	// password := os.Getenv("MONGO_PASSWORD")

	if os.Getenv("MONGO_URI") == "" {
		logger.Info("Starting mongoDB connection to mongodb://localhost:27017")
		return options.Client().ApplyURI("mongodb://localhost:27017")
	}

	logger.GetLogger().Printf("Start MongoDB connection to %s", os.Getenv("MONGO_URI"))
	return options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
}
