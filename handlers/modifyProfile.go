package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.UserInfo

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error request body inccorect "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = db.ModifyRecord(t, IDUserInfo)
	if err != nil {
		http.Error(w, "An error occurred while modify register "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "It was not possible to modify the user registry "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
