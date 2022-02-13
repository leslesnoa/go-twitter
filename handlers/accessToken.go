package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt"
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

var (
	Email      string
	IDUserInfo string
)

func AccessToken(tk string, ctx context.Context) (*models.Claim, bool, string, error) {
	signKey := []byte(os.Getenv("SIGN_KEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return nil, false, string(""), errors.New(fmt.Sprintf("format error invalid token %s", tk))
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	if err != nil {
		return nil, false, string(""), err
	}

	if !tkn.Valid {
		return nil, false, string(""), errors.New("token invalid")
	}

	_, isExist, _ := db.CheckIsExistUser(claims.Email, ctx)
	if isExist == true {
		Email = claims.Email
		IDUserInfo = claims.ID.Hex()
	}

	return claims, isExist, IDUserInfo, nil
}
