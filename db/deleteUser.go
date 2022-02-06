package db

import (
	"context"
	"time"

	"github.com/leslesnoa/go-twitter/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		logger.Error("Error while delete user operation", err)
	}
	return err
}
