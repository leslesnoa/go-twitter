package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leslesnoa/go-twitter/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error You must send the id parameter ", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Error You must send the page parameter ", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Error You must send the page parameter with a value greater than 0", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	pag := int64(page)
	response, right := db.ReadTweets(ID, pag, ctx)
	if right == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
