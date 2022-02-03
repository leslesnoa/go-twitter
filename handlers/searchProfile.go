package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
)

func SearchProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "parameter ID error", http.StatusBadRequest)
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "an error occurred while search record "+err.Error(), 400)
	}

	w.Header().Set("contnt-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
