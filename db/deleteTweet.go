package db

import (
	"context"
	"time"

	"github.com/leslesnoa/go-twitter/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

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
