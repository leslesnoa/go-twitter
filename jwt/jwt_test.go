package jwt

import (
	"testing"

	jwt "github.com/golang-jwt/jwt"
	"github.com/leslesnoa/go-twitter/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGenerateJWT(t *testing.T) {

	signKey = []byte("sign key by test")
	objID := primitive.NewObjectID()

	u := models.UserInfo{
		Email: "test@gmail.com",
		ID:    objID,
	}

	jwtKey, err := GenerateJWT(&u)
	// os.Setenv("testToken", jwtKey)

	/* 正常にJWT生成がされること */
	assert.NoError(t, err)

	claims := &models.Claim{}
	tk, err := jwt.ParseWithClaims(jwtKey, claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	/* 正常にtokenをパースできること */
	assert.NoError(t, err)

	/* Tokenが有効であること */
	assert.Equal(t, tk.Valid, true)

	/* claimの情報が正しいこと */
	assert.Equal(t, claims.Email, "test@gmail.com")
	assert.Equal(t, claims.ID, objID)
}
