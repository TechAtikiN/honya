package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/techatikin/backend/errors"
)

type EnvConfig struct {
	DatabaseURL              string
	ServerPort               string
	LogStack                 string
	LogRetention             string
	UrlCleanupOriginalDomain string
	AWSBucket                string
	AWSRegion                string
	AWSAccessKey             string
	AWSSecretKey             string
}

var NewEnvConfig struct {
	DatabaseURL              string
	ServerPort               string
	LogStack                 string
	LogRetention             string
	UrlCleanupOriginalDomain string
	AWSBucket                string
	AWSRegion                string
	AWSAccessKey             string
	AWSSecretKey             string
}

func GetEnvConfig() (EnvConfig, error) {
	if _, err := os.Stat(".env.local"); err == nil {
		err := godotenv.Load(".env.local", ".env")

		if err != nil {
			return NewEnvConfig, errors.NewBadRequestError("Error loading .env.local file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			return NewEnvConfig, errors.NewBadRequestError("Error loading .env file")
		}
	}

	NewEnvConfig.DatabaseURL = os.Getenv("DATABASE_URL")
	if NewEnvConfig.DatabaseURL == "" {
		return NewEnvConfig, errors.NewBadRequestError("DATABASE_URL environment variable is not set")
	}

	NewEnvConfig.ServerPort = os.Getenv("SERVER_PORT")
	if NewEnvConfig.ServerPort == "" {
		NewEnvConfig.ServerPort = "8080"
	}

	NewEnvConfig.LogStack = os.Getenv("LOG_STACK")
	if NewEnvConfig.LogStack == "" {
		NewEnvConfig.LogStack = "daily"
	}

	NewEnvConfig.LogRetention = os.Getenv("LOG_RETENTION")
	if NewEnvConfig.LogRetention == "" {
		NewEnvConfig.LogRetention = "7"
	}

	NewEnvConfig.UrlCleanupOriginalDomain = os.Getenv("URL_CLEANUP_ORIGINAL_DOMAIN")
	if NewEnvConfig.UrlCleanupOriginalDomain == "" {
		return NewEnvConfig, errors.NewBadRequestError("URL_CLEANUP_ORIGINAL_DOMAIN environment variable is not set")
	}

	NewEnvConfig.AWSBucket = os.Getenv("AWS_BUCKET_NAME")
	if NewEnvConfig.AWSBucket == "" {
		return NewEnvConfig, errors.NewBadRequestError("AWS_BUCKET_NAME environment variable is not set")
	}

	NewEnvConfig.AWSRegion = os.Getenv("AWS_REGION")
	if NewEnvConfig.AWSRegion == "" {
		return NewEnvConfig, errors.NewBadRequestError("AWS_REGION environment variable is not set")
	}

	NewEnvConfig.AWSAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	if NewEnvConfig.AWSAccessKey == "" {
		return NewEnvConfig, errors.NewBadRequestError("AWS_ACCESS_KEY_ID environment variable is not set")
	}

	NewEnvConfig.AWSSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if NewEnvConfig.AWSSecretKey == "" {
		return NewEnvConfig, errors.NewBadRequestError("AWS_SECRET_ACCESS_KEY environment variable is not set")
	}

	return NewEnvConfig, nil
}
