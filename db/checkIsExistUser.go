package db

import (
	"context"
	"time"

	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIsExistUser(email string) (models.UserInfo, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resUser models.UserInfo

	err := col.FindOne(ctx, condicion).Decode(&resUser)
	ID := resUser.ID.Hex()
	if err != nil {
		return resUser, false, ID
	}
	return resUser, true, ID
}
