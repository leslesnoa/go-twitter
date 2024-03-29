package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.UserInfo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Error invalid request Email is cannot empty", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 4 {
		http.Error(w, "Error password cannot less than 6 characters", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	_, encontrado, _ := db.CheckIsExistUser(t.Email, ctx)
	if encontrado == true {
		http.Error(w, "Error invalid request Email is already registerd", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(t, ctx)
	if err != nil {
		http.Error(w, "An error ocurred while register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "an error occured insert into user record", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
