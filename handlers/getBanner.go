package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/logger"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "must request parameter id", http.StatusBadRequest)
		logger.GetLogger().Printf("Error request parameter id not include %s", ID)
		return
	}

	ctx := r.Context()

	profile, err := db.SearchProfile(ID, ctx)
	if err != nil {
		http.Error(w, "must request parameter id "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while SearchProfile ", err)
		return
	}

	Openfile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Error not Found an image "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while openfile by request user image ", err)
		return
	}

	_, err = io.Copy(w, Openfile)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while copyfile by request user image ", err)
		return
	}

	defer Openfile.Close()
}
