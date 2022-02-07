package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.PostTweet, ctx context.Context) (string, bool, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	register := bson.M{
		"user_id": t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		logger.Error("Error while insert tweet", err)
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
