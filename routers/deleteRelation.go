package routers

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

	status, err := db.DeleteRelation(t)
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
