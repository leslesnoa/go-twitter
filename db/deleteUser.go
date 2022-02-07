package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(ID string, ctx context.Context) error {

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
