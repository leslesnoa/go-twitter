package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string, ctx context.Context) error {

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":     objID,
		"user_id": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		logger.Error("Error while Delete tweet", err)
	}
	return err
}
