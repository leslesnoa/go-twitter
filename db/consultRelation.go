package db

import (
	"context"
	"fmt"
	"time"

	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	condition := bson.M{
		"user_id":          t.UserID,
		"user_relation_id": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
