package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leslesnoa/go-twitter/db"
	jwt "github.com/leslesnoa/go-twitter/jwt"
	"github.com/leslesnoa/go-twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.UserInfo

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid request login user: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "invalid request Email: "+err.Error(), 400)
		return
	}

	document, isExist := db.TryLogin(t.Email, t.Password)
	if isExist == false {
		http.Error(w, "invalid username and / or password "+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred when trying to generate the corresponding Token: "+err.Error(), 400)
		return
	}

	resp := models.LoginRequest{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
