package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string, ctx context.Context) (models.UserInfo, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var profile models.UserInfo
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		logger.Error("An error occured while show user profile", err)
		return profile, err
	}
	return profile, nil
}
