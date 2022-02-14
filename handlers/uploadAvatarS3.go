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

var (
	// 画像のPath
	filePath = "../testimage.jpeg"
	// S3のバケット名
	bucketName = "test-bucket0215"
	// key S3に保存するオブジェクトの名前になります
	// key = "image/avatar"
	// awsのリージョン名
	awsRegion = "ap-northeast-1"
)

func UploadAvatarS3(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	defer file.Close()

	var extention = strings.Split(handler.Filename, ".")[1]
	var key string = "uploads/avatars/" + IDUserInfo + "." + extention

	// f, err := os.Open(file)
	// if err != nil {
	// 	http.Error(w, "Error when uploading image! "+err.Error(), http.StatusBadRequest)
	// 	logger.Error("An Error ocurred while open file ", err)
	// 	return
	// }
	// defer f.Close()

	// セッションとデフォルトのオプションを使用してアップローダーを作成
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: "default",
	})
	if err != nil {
		http.Error(w, "Error when uploading image! "+err.Error(), http.StatusBadRequest)
		logger.Error("An Error ocurred while new session ", err)
		return
	}

	uploader := s3manager.NewUploader(sess)

	// f, err := os.Open(key)
	// if err != nil {
	// 	http.Error(w, "Error when uploading image! "+err.Error(), http.StatusBadRequest)
	// 	logger.Error("An Error ocurred while open file ", err)
	// 	return
	// }

	// _, err = io.Copy(f, file)
	// if err != nil {
	// 	http.Error(w, "Error when copy image! "+err.Error(), http.StatusBadRequest)
	// 	logger.Error("An Error while copy file ", err)
	// 	return
	// }

	// ファイルをS3にアップロードする
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		http.Error(w, "Error when copy image! "+err.Error(), http.StatusBadRequest)
		logger.Error("An Error while uploading file ", err)
		// return fmt.Errorf("failed to upload file, %v", err)
	}

	logger.Info("file upload success")

	var user models.UserInfo
	var status bool
	ctx := r.Context()

	user.Avatar = IDUserInfo + "." + extention
	status, err = db.ModifyRecord(user, IDUserInfo, ctx)
	if err != nil || status == false {
		http.Error(w, "Error when saving the avatar in the DB! "+err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Printf("file uploaded to, %s\n", aws.StringValue(result.Location))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
