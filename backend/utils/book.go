package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/techatikin/backend/dto"
)

func ValidateBookCreateRequest(request *dto.BookCreateRequest) error {
	if request.Title == "" {
		return errors.New("title is required")
	}
	if request.AuthorName == "" {
		return errors.New("author ID is required")
	}
	currentYear := time.Now().Year()
	if request.PublicationYear < 1950 || request.PublicationYear > currentYear {
		return errors.New("publication year must be between 1950 and " + strconv.Itoa(currentYear))
	}
	if request.Rating < 0 || request.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}
	if request.Pages <= 0 {
		return errors.New("pages must be a positive integer")
	}
	if request.Isbn == "" {
		return errors.New("ISBN is required")
	}
	return nil
}

func ValidateBookUpdateRequest(request *dto.BookUpdateRequest) error {
	currentYear := time.Now().Year()

	if request.Title != nil && *request.Title == "" {
		return errors.New("title cannot be empty")
	}
	if request.AuthorName != nil && *request.AuthorName == "" {
		return errors.New("author name cannot be empty")
	}
	if request.PublicationYear != nil {
		if *request.PublicationYear < 1950 || *request.PublicationYear > currentYear {
			return errors.New("publication year must be between 1950 and " + strconv.Itoa(currentYear))
		}
	}
	if request.Rating != nil {
		if *request.Rating < 0 || *request.Rating > 5 {
			return errors.New("rating must be between 0 and 5")
		}
	}
	if request.Pages != nil && *request.Pages <= 0 {
		return errors.New("pages must be a positive integer")
	}

	return nil
}

func ExtractS3Key(url, bucket, region string) string {
	prefix := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucket, region)
	return strings.TrimPrefix(url, prefix)
}
