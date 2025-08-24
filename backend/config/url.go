package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/techatikin/backend/errors"
)

func GetOriginalDomain() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.NewBadRequestError("Error loading .env file")
	}

	domain := os.Getenv("URL_CLEANUP_ORIGINAL_DOMAIN")
	if domain == "" {
		domain = "www.example.com"
	}

	return domain, nil
}
