package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRecord(u models.UserInfo, ID string, ctx context.Context) (bool, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	// make関数でJSONを表現
	register := make(map[string]interface{})
	if len(u.Number) > 0 {
		register["number"] = u.Number
	}
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Birth) > 0 {
		register["birth"] = u.Birth
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Biografia) > 0 {
		register["biografia"] = u.Biografia
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.Name) > 0 {
		register["website"] = u.WebSite
	}

	updtString := bson.M{
		"$set": register,
	}

	/* 渡されたIDに等しいユーザを抽出するフィルター */
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Error("Error request ID is invalid", err)
		return false, err
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err = col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		logger.Error("Error while update user profile", err)
		return false, err
	}

	return true, nil
}
