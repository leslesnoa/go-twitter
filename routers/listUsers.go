package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leslesnoa/go-twitter/db"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("kind")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "you must page parameter request send", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUser(IDUserInfo, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error while reading users DB", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
