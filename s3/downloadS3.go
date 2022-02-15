package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/leslesnoa/go-twitter/logger"
)

func DownloadS3(fileName string) (*os.File, error) {
	key := uploadAvatarPath + fileName

	// S3オブジェクトのコンテンツを書き込むファイルを作成
	f, err := os.Create(key)
	if err != nil {
		logger.Error("Error while create image ", err)
		return nil, err
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
			Region:      aws.String(awsRegion),
		},
		Profile: "default",
	})
	if err != nil {
		logger.Error("Error could not create new session ", err)
		return nil, err
	}

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// S3オブジェクトの内容をファイルに書き込む
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		logger.Error("Error could not download image from S3 ", err)
		return nil, err
	}

	return f, nil
}
