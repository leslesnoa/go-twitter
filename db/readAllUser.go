package db

import (
	"context"
	"fmt"
	"time"

	"github.com/leslesnoa/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUser(ID string, page int64, search string, kind string) ([]*models.UserInfo, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var results []*models.UserInfo

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"number": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var relation, include bool

	for cur.Next(ctx) {
		var s models.UserInfo
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		relation, err = ConsultRelation(r)
		if kind == "new" && relation == false {
			include = true
		}

		if kind == "follow" && relation == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.WebSite = ""
			s.Email = ""
			s.Location = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
