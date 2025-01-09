package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gabriel-vasile/mimetype"
)

// CreateSession creates a new AWS session
func CreateSession() (*session.Session, error) {
	region := os.Getenv("AWS_REGION")

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})
}

// UploadFileToS3 uploads a file to S3 with a given prefix
func UploadFileToS3(filePath *multipart.FileHeader, bucket, prefix string) (string, error) {
	sess, err := CreateSession()
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	file, err := filePath.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}

	/// Reset the file reader to the beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Read the file into a buffer to determine its size
	var buf bytes.Buffer
	fileSize, err := io.Copy(&buf, file)
	if err != nil {
		return "", err
	}

	// Reset the file reader to the beginning again
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	timestamp := time.Now().Format("20060102150405")
	objectKey := fmt.Sprintf("%s/%s_%s", prefix, timestamp, filePath.Filename)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(objectKey),
		Body:          file,
		ContentType:   aws.String(mime.String()),
		ContentLength: aws.Int64(fileSize),
	})

	region := os.Getenv("AWS_REGION")
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, objectKey)

	return url, err
}
