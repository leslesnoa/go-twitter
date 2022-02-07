package handlers

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
		http.Error(w, "invalid request body for login user: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "invalid request Email: "+err.Error(), http.StatusBadRequest)
		return
	}

	document, isExist := db.TryLogin(t.Email, t.Password)
	if !isExist {
		http.Error(w, "invalid request username or password", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred when trying to generate the corresponding Token: "+err.Error(), http.StatusBadRequest)
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
