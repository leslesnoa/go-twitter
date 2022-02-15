package handlers

import (
	"net/http"
	"strings"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
	"github.com/leslesnoa/go-twitter/s3"
)

const (
	uploadAvatarPath = "uploads/avatars/"
	bucketName       = "test-bucket0215"
	awsRegion        = "ap-northeast-1"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	defer file.Close()

	var extention = strings.Split(handler.Filename, ".")[1]
	var key string = uploadAvatarPath + IDUserInfo + "." + extention

	if err := s3.UploadS3(file, key); err != nil {
		http.Error(w, "Error when request image uploading S3! "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.UserInfo
	var status bool
	ctx := r.Context()

	user.Avatar = IDUserInfo + "." + extention
	status, err = db.ModifyRecord(user, IDUserInfo, ctx)
	if err != nil || status == false {
		http.Error(w, "Error when saving the avatar in the DB! "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
