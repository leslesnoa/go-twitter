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
		http.Error(w, "Error You must send the id parameter ", 400)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Error You must send the page parameter ", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Error You must send the page parameter with a value greater than 0", 400)
		return
	}

	pag := int64(page)
	response, right := db.ReadTweets(ID, pag)
	if right == false {
		http.Error(w, "Error reading tweets", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
