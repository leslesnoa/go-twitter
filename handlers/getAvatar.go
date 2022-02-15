package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/logger"
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
		http.Error(w, "Error request must into id "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while SearchProfile ", err)
		return
	}

	key := uploadAvatarPath + profile.Avatar

	// S3オブジェクトのコンテンツを書き込むファイルを作成
	f, err := os.Create(key)
	if err != nil {
		http.Error(w, "Error could not create image "+err.Error(), http.StatusInternalServerError)
		logger.Error("Error while create image ", err)
		return
	}
	defer f.Close()

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: "default",
	})
	if err != nil {
		http.Error(w, "Error could not get avatar "+err.Error(), http.StatusInternalServerError)
		logger.Error("Error could not create new session ", err)
		return
	}

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// S3オブジェクトの内容をファイルに書き込みます
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		http.Error(w, "Error could not get avatar "+err.Error(), http.StatusInternalServerError)
		logger.Error("Error could not download image from S3 ", err)
		return
	}

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusInternalServerError)
		logger.Error("Error while copyfile by request user image ", err)
		return
	}
}
