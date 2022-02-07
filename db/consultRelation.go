package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultRelation(t models.Relation, ctx context.Context) (bool, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	condition := bson.M{
		"user_id":          t.UserID,
		"user_relation_id": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		// logger.Error("Error while consult relation", err)
		return false, err
	}
	return true, nil
}
