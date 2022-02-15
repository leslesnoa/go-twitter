package handlers

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
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

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: "default",
	})
	if err != nil {
		http.Error(w, "Error when uploading image! "+err.Error(), http.StatusInternalServerError)
		logger.Error("An Error ocurred while new session ", err)
		return
	}

	uploader := s3manager.NewUploader(sess)

	// requestファイルをS3にアップロードする
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		http.Error(w, "Error when copy image! "+err.Error(), http.StatusInternalServerError)
		logger.Error("An Error while uploading file ", err)
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
