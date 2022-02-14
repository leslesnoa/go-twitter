package handlers

import (
	"fmt"
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

func GetAvatarS3(w http.ResponseWriter, r *http.Request) {
	// The session the S3 Downloader will use
	// sess := session.Must(session.NewSession())
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
	// ------------------
	key := "uploads/avatars/" + profile.Avatar
	logger.Info(key)
	logger.Info(profile.Avatar)
	f, err := os.Create(key)
	if err != nil {
		http.Error(w, "Error could not create image "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while create image ", err)
		return
	}
	defer f.Close()

	// ------------------

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: "default",
	})
	if err != nil {
		http.Error(w, "Error could not get avatar "+err.Error(), http.StatusBadRequest)
		logger.Error("Error could not create new session ", err)
		return
	}
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// S3オブジェクトのコンテンツを書き込むファイルを作成します。
	// f, err := os.Create("gorira.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// return fmt.Errorf("failed to create file %q, %v", filename, err)
	// }
	// defer f.Close()

	// S3オブジェクトの内容をファイルに書き込みます
	res, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		http.Error(w, "Error could not get avatar "+err.Error(), http.StatusBadRequest)
		logger.Error("Error could not download image from S3 ", err)
		return
		// return fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Println(res)
	// os.Open(key)

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Error when copy image "+err.Error(), http.StatusBadRequest)
		logger.Error("Error while copyfile by request user image ", err)
		return
	}
	logger.Info("file download successfuly!")
	// fmt.Printf("file downloaded, %d bytes\n", n)
}
