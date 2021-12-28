package routers

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
		http.Error(w, "error data inccorect "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyRecord(t, IDUserInfo)
	if err != nil {
		http.Error(w, "Occurred an error while modify register "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It was not possible to modify the user registry "+err.Error(), 400)
	}

	w.WriteHeader(http.StatusCreated)

}
