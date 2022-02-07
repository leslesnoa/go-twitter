package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64, ctx context.Context) ([]*models.TweetsResponse, bool) {

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var results []*models.TweetsResponse

	condition := bson.M{
		"user_id": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	/*日付の降順でソート*/
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		logger.Error("Error while read tweets", err)
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.TweetsResponse
		if err := cursor.Decode(&register); err != nil {
			logger.Error("Error while read tweets", err)
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}
