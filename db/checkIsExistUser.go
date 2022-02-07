package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIsExistUser(email string, ctx context.Context) (*models.UserInfo, bool, string) {

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resUser models.UserInfo

	err := col.FindOne(ctx, condicion).Decode(&resUser)
	ID := resUser.ID.Hex()
	if err != nil {
		return nil, false, ID
	}
	return &resUser, true, ID
}
