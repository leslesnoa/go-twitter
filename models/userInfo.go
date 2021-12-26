package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Number string             `bson:"number" json:"number,omitempty"`
	Name   string             `bson:"name" json:"name,omitempty"`
	Birth  string             `bson:"birth" json:"birth,omitempty"`
	// Birth     time.Time          `bson:"birth" json:"birth,omitempty"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password,omitempty"`
	Avatar    string `bson:"avatar" json:"avatar,omitempty"`
	Banner    string `bson:"banner" json:"banner,omitempty"`
	Biografia string `bson:"biografia" json:"biografia,omitempty"`
	Location  string `bson:"location" json:"location,omitempty"`
	WebSite   string `bson:"website" json:"website,omitempty"`
}
