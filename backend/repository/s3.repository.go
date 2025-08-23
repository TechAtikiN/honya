package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Repository interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
	DeleteImage(key string) error
}

type s3Repository struct {
	client     *s3.Client
	bucketName string
	region     string
}

func NewS3Repository() S3Repository {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)

	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(cfg)
	return &s3Repository{client: client, bucketName: bucket, region: region}
}

func (r *s3Repository) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := fmt.Sprintf("books/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	uploader := manager.NewUploader(r.client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.bucketName, r.region, key)
	return url, nil
}

func (r *s3Repository) DeleteImage(key string) error {
	_, err := r.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(key),
	})
	return err
}
