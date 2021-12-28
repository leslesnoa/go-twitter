package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TweetsResponse struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"user_id" json:"user_id,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    string             `bson:"date" json:"date,omitempty"`
}
