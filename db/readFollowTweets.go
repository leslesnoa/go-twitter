package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadFllowTweets(ID string, page int, ctx context.Context) ([]models.ResponseFollowerTweets, bool) {

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"user_id": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "user_relation_id",
			"foreignField": "user_id",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)

	var results []models.ResponseFollowerTweets
	err = cursor.All(ctx, &results)
	if err != nil {
		logger.Error("Error while read follow tweets", err)
		return results, false
	}
	return results, true
}
