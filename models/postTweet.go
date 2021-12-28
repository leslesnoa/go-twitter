package models

type PostTweet struct {
	UserId  string `bson:"user_id" json:"user_id,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date    string `bson:"date" json:"date,omitempty"`
}
