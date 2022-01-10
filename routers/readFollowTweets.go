package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leslesnoa/go-twitter/db"
)

func ReadFollowTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must request page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Can not page parameter is 0", http.StatusBadRequest)
		return
	}

	request, ok := db.ReadFllowTweets(IDUserInfo, page)
	if ok == false {
		http.Error(w, "Error while read tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}
