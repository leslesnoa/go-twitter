package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/leslesnoa/go-twitter/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error request must be parameter id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	profile, err := db.SearchProfile(ID, ctx)
	if err != nil {
		http.Error(w, "Error request must into id "+err.Error(), http.StatusBadRequest)
		return
	}

	Openfile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Error not Found an image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Openfile)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusBadRequest)
		return
	}
	defer Openfile.Close()
}
