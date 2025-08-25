package dto

import (
	"honya/backend/model"

	"github.com/google/uuid"
)

// Request payload for creating a review
type ReviewCreateRequest struct {
	BookID  uuid.UUID `json:"book_id" validate:"required"`
	Name    string    `json:"name" validate:"required"`
	Email   string    `json:"email" validate:"required,email"`
	Content string    `json:"content" validate:"required"`
}

// Request payload for updating a review
type ReviewUpdateRequest struct {
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty" validate:"omitempty,email"`
	Content *string `json:"content,omitempty"`
}

// Response payload for a single review
type ReviewResponse struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

// Response for list of reviews
type ReviewListResponse struct {
	Meta PaginationMeta   `json:"meta"`
	Data []ReviewResponse `json:"data"`
}

// Convert Review model -> ReviewResponse
func ToReviewResponse(review *model.Review) *ReviewResponse {
	return &ReviewResponse{
		ID:        review.ID,
		BookID:    review.BookID,
		Name:      review.Name,
		Email:     review.Email,
		Content:   review.Content,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}

// Convert slice of Reviews -> ReviewListResponse
func ToReviewListResponse(reviews []model.Review, meta PaginationMeta) ReviewListResponse {
	responses := make([]ReviewResponse, 0, len(reviews))
	for _, r := range reviews {
		responses = append(responses, *ToReviewResponse(&r))
	}

	return ReviewListResponse{
		Meta: meta,
		Data: responses,
	}
}
