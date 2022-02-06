package handlers

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Error invalid parameter id must be in query parameter", http.StatusBadRequest)
		return
	}

	/* ID: objID, IDUserInfo: UserID */
	err := db.DeleteUser(IDUserInfo)
	if err != nil {
		http.Error(w, "Error while delete tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
