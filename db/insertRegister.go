package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u models.UserInfo, ctx context.Context) (string, bool, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		logger.Error("Error while registering user", err)
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.Hex(), true, nil
}
