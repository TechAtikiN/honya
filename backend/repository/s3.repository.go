package repository

import (
	"context"
	"fmt"
	"honya/backend/config"
	"honya/backend/utils"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Repository interface {
	UploadImage(fileHeader *multipart.FileHeader, bookName string) (string, error)
	DeleteImage(key string) error
}

type s3Repository struct {
	client     *s3.Client
	bucketName string
	region     string
}

func NewS3Repository() S3Repository {
	env, err := config.GetEnvConfig()
	if err != nil {
		fmt.Println("Failed to get environment configuration:", err)
		return nil
	}

	AWS_BUCKET := env.AWSBucket
	AWS_REGION := env.AWSRegion
	AWS_ACCESS_KEY := env.AWSAccessKey
	AWS_SECRET_KEY := env.AWSSecretKey

	cfg, err := s3_config.LoadDefaultConfig(context.TODO(),
		s3_config.WithRegion(AWS_REGION),
		s3_config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
		),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(cfg)
	return &s3Repository{client: client, bucketName: AWS_BUCKET, region: AWS_REGION}
}

func (r *s3Repository) UploadImage(fileHeader *multipart.FileHeader, bookName string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create safe key
	bookSlug := utils.Slugify(bookName)
	key := fmt.Sprintf("books/%s-%d%s", bookSlug, time.Now().Unix(), filepath.Ext(fileHeader.Filename))

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
