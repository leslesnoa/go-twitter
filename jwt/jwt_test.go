package jwt_test

import (
	"os"
	"testing"

	jwtt "github.com/dgrijalva/jwt-go"
	"github.com/leslesnoa/go-twitter/jwt"
	"github.com/leslesnoa/go-twitter/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	signKey = []byte(os.Getenv("SIGN_KEY"))
)

func TestGenerateJWT(t *testing.T) {

	objID := primitive.NewObjectID()

	u := models.UserInfo{
		Email: "test@gmail.com",
		ID:    objID,
	}

	jwtKey, err := jwt.GenerateJWT(u)
	// os.Setenv("testToken", jwtKey)

	/* 正常にJWT生成がされること */
	assert.NoError(t, err)

	claims := &models.Claim{}
	tk, err := jwtt.ParseWithClaims(jwtKey, claims, func(token *jwtt.Token) (interface{}, error) {
		return signKey, nil
	})

	/* 正常にtokenをパースできること */
	assert.NoError(t, err)

	/* Tokenが有効であること */
	assert.Equal(t, tk.Valid, true)

	/* claimの情報が正しいこと */
	assert.Equal(t, claims.Email, "test@gmail.com")
	assert.Equal(t, claims.ID, objID)
	assert.Equal(t, claims.ExpiresAt, objID)
}
