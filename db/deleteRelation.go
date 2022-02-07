package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
)

func DeleteRelation(t models.Relation, ctx context.Context) (bool, error) {

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		logger.Error("Error while delete reration", err)
		return false, err
	}
	return true, nil
}
