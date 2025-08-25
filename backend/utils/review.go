package utils

import (
	"errors"
	"honya/backend/dto"
	"net/mail"

	"github.com/google/uuid"
)

func ValidateReviewCreateRequest(request *dto.ReviewCreateRequest) error {
	if request.BookID == uuid.Nil {
		return errors.New("book_id is required")
	}
	if request.Name == "" {
		return errors.New("name is required")
	}
	if request.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		return errors.New("invalid email format")
	}
	if request.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func ValidateReviewUpdateRequest(request *dto.ReviewUpdateRequest) error {
	if request.Name != nil && *request.Name == "" {
		return errors.New("name cannot be empty")
	}
	if request.Email != nil {
		if *request.Email == "" {
			return errors.New("email cannot be empty")
		}
		if _, err := mail.ParseAddress(*request.Email); err != nil {
			return errors.New("invalid email format")
		}
	}
	if request.Content != nil && *request.Content == "" {
		return errors.New("content cannot be empty")
	}
	return nil
}
