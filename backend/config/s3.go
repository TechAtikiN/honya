package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/techatikin/backend/errors"
)

func GetS3Config() (string, string, string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", "", "", errors.NewBadRequestError("Error loading .env file")
	}

	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if bucket == "" || region == "" || accessKey == "" || secretKey == "" {
		return "", "", "", "", errors.NewBadRequestError("Missing S3 configuration in environment variables")
	}

	return bucket, region, accessKey, secretKey, nil
}
