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
		http.Error(w, "must request parameter id", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "must request parameter id "+err.Error(), http.StatusBadRequest)
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

	if err := Openfile.Close(); err != nil {
		http.Error(w, "Internal Server Error image could not closed "+err.Error(), http.StatusInternalServerError)
		return
	}
}
