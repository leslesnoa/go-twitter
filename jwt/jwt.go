package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
)

func GenerateJWT(t models.UserInfo) (string, error) {

	signKey := []byte(os.Getenv("SIGN_KEY"))
	// signKey := []byte("DevelopmentMasters_Facebookgroup")
	logger.Info(string(signKey))

	claims := jwt.MapClaims{
		"email":     t.Email,
		"number":    t.Number,
		"name":      t.Name,
		"birth":     t.Birth,
		"biografia": t.Biografia,
		"location":  t.Location,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(signKey))
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
