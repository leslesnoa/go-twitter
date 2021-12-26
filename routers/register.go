package routers

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
		http.Error(w, "invalid your request body: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is cannot empty: "+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password cannot less than 6 characters: "+err.Error(), 400)
		return
	}

	_, encontrado, _ := db.CheckIsExistUser(t.Email)
	if encontrado == true {
		http.Error(w, "request Email is already registerd", 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Occured un error while register user: "+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "an error occured insert into user record", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
