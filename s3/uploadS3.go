package s3

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/leslesnoa/go-twitter/logger"
)

const (
	uploadAvatarPath = "uploads/avatars/"
	bucketName       = "test-bucket0215"
	awsRegion        = "ap-northeast-1"
)

var (
	awsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey = os.Getenv("AWS_SECRET_KEY")
)

func UploadS3(f multipart.File, key string) error {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
			Region:      aws.String(awsRegion),
		},
		Profile: "default",
	})
	if err != nil {
		logger.Error("Internal Server Error ocurred while create new session ", err)
		return err
	}

	uploader := s3manager.NewUploader(sess)

	// requestファイルをS3にアップロードする
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		logger.Error("Error while uploading file to the S3 ", err)
		return err
	}

	return nil
}
