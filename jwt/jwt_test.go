package jwt_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/leslesnoa/go-twitter/jwt"
	"github.com/leslesnoa/go-twitter/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* 何もしないhttp.HandlerFuncをHTTP Middlewareに渡す */
func getTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test Handler")
	}
	return http.HandlerFunc(fn)
}

/* それをhttptest.NewServerに渡してテストサーバを立ち上げる */
// func TestGenerateJWT(t *testing.T) {
// 	ts := httptest.NewServer(middleware.ValidJWT(getTestHandler()))
// 	defer ts.Close()

// 	resp, err := http.Get(ts.URL)
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	fmt.Println(resp.Body)
// 	// assert.Equal()
// }

// func GetTestHandler() http.HandlerFunc {
// 	fn := func(rw http.ResponseWriter, req *http.Request) {
// 		rw.Write([]byte("OK"))
// 		return
// 	}
// 	return http.HandlerFunc(fn)
// }

func TestGenerateJWT(t *testing.T) {
	// jwtWrapper := JwtWrapper{
	// 	SecretKey:       "verysecretkey",
	// 	Issuer:          "AuthService",
	// 	ExpirationHours: 24,
	// }

	jwtKey, err := jwt.GenerateJWT(models.UserInfo{
		Email:     "test@gmail.com",
		Number:    "田中",
		Name:      "太郎",
		Birth:     "2000-12-12",
		Biografia: "test profile",
		Location:  "test location",
		WebSite:   "test.com",
		ID:        primitive.NewObjectID(),
	})
	assert.NoError(t, err)

	os.Setenv("testToken", jwtKey)
}

func TestAccessToken(t *testing.T) {

}

// func TestValidateToken(t *testing.T) {
// 	encodedToken := os.Getenv("testToken")

// 	// jwtWrapper := JwtWrapper{
// 	// 	SecretKey: "verysecretkey",
// 	// 	Issuer:    "AuthService",
// 	// }
// 	ts := httptest.NewServer(middleware.ValidJWT(GetTestHandler()))
// 	defer ts.Close()
// 	var u bytes.Buffer
// 	u.WriteString(string(ts.URL))
// 	u.WriteString(tt.pass)

// 	req, _ := http.NewRequest("GET", u.String(), nil)
// 	// claims, err := middleware.ValidJWT(encodedToken)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "jwt@email.com", claims.Email)
// 	assert.Equal(t, "AuthService", claims.Issuer)
// }
