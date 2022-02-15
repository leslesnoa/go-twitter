package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/logger"
	s3pkg "github.com/leslesnoa/go-twitter/s3"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error request must be parameter id", http.StatusBadRequest)
		logger.GetLogger().Printf("Error request parameter id not include %s", ID)
		return
	}

	ctx := r.Context()
	profile, err := db.SearchProfile(ID, ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error invalid request id: %s ", ID)+err.Error(), http.StatusBadRequest)
		logger.Error("Error while SearchProfile ", err)
		return
	}

	f, err := s3pkg.DownloadS3(profile.Avatar)
	if err != nil {
		http.Error(w, "Internal Server Error ocuured while downloading file for S3 "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusInternalServerError)
		logger.Error("Error while copyfile by request user image ", err)
		return
	}
}
