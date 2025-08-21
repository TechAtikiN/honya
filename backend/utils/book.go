package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/techatikin/backend/dtos"
)

func ValidateBookCreateRequest(request dtos.BookCreateRequest) error {
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

func ValidateBookUpdateRequest(request dtos.BookUpdateRequest) error {
	if *request.Title == "" {
		return errors.New("title is required")
	}
	if *request.AuthorName == "" {
		return errors.New("author ID is required")
	}
	currentYear := time.Now().Year()
	if *request.PublicationYear < 1950 || *request.PublicationYear > currentYear {
		return errors.New("publication year must be between 1950 and " + strconv.Itoa(currentYear))
	}
	if *request.Rating < 0 || *request.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}
	if *request.Pages <= 0 {
		return errors.New("pages must be a positive integer")
	}
	if *request.Isbn == "" {
		return errors.New("ISBN is required")
	}
	return nil
}
