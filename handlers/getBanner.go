package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/leslesnoa/go-twitter/db"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "must request parameter id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	profile, err := db.SearchProfile(ID, ctx)
	if err != nil {
		http.Error(w, "must request parameter id "+err.Error(), http.StatusBadRequest)
		return
	}

	Openfile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Error not Found an image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Openfile)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusBadRequest)
		return
	}

}
