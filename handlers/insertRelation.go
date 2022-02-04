package handlers

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func Relation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "invalid parameter ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUserInfo
	t.UserRelationID = ID

	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "an error occured while insert relation "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "invalid parameter ID", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
