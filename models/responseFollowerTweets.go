package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseFollowerTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"user_id" json:"user_id,omitempty"`
	UserRelationID string             `bson:"user_relation_id" json:"user_relation_id,omitempty"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
