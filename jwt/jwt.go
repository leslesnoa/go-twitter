package jwt

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/leslesnoa/go-twitter/models"
)

func GenerateJWT(t models.UserInfo) (string, error) {

	signKey := []byte("DevelopmentMasters_Facebookgroup")

	// payload := jwt.MapClaims{
	// 	"email":     t.Email,
	// 	"number":    t.Number,
	// 	"name":      t.Name,
	// 	"birth":     t.Birth,
	// 	"biografia": t.Biografia,
	// 	"location":  t.Location,
	// 	"website":   t.WebSite,
	// 	"_id":       t.ID.Hex(),
	// 	"exp":       time.Now().Add(time.Hour * 24).Unix(),
	// }
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = t.Name
	claims["number"] = t.Email
	claims["name"] = t.Name
	claims["birth"] = t.Birth
	claims["biografia"] = t.Biografia
	claims["location"] = t.Location
	claims["website"] = t.WebSite
	claims["_id"] = t.ID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString([]byte(signKey))
	if err != nil {
		log.Println("署名時エラー")
		return tokenStr, err
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	// tokenStr, err := token.SignedString(miClave)
	// if err != nil {
	// 	return tokenStr, err
	// }

	return tokenStr, nil
}
