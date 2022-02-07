package handlers

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUserInfo
	t.UserRelationID = ID

	ctx := r.Context()

	status, err := db.DeleteRelation(t, ctx)
	if err != nil {
		http.Error(w, "an error occured while delete relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		if err != nil {
			http.Error(w, "an error occured while delete relation "+err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
