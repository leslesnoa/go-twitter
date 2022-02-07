package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func PostTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.PostTweet{
		UserId:  IDUserInfo,
		Message: message.Message,
		Date:    time.Now(),
	}

	ctx := r.Context()

	_, status, err := db.InsertTweet(register, ctx)
	if err != nil {
		http.Error(w, "Occured an error while insert tweet register "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Insert failed "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
