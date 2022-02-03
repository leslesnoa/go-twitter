package handlers

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

var (
	Email      string
	IDUserInfo string
)

func AccessToken(tk string) (*models.Claim, bool, string, error) {
	signKey := []byte(os.Getenv("SIGN_KEY"))
	// signKey := []byte("DevelopmentMasters_Facebookgroup")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("format error invalid token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err == nil {
		_, isExist, _ := db.CheckIsExistUser(claims.Email)
		if isExist == true {
			Email = claims.Email
			IDUserInfo = claims.ID.Hex()
		}
		return claims, isExist, IDUserInfo, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalid")
	}

	return claims, false, string(""), err
}
