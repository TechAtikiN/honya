package dtos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/techatikin/backend/models"
)

type BookQueryParams struct {
	Offset          int     `query:"offset"`
	Limit           int     `query:"limit"`
	Query           string  `query:"query"`
	Category        string  `query:"category"`
	PublicationYear int     `query:"publication_year"`
	Rating          float64 `query:"rating"`
	Pages           int     `query:"pages"`
	Sort            string  `query:"sort"`
}

type PaginationMeta struct {
	TotalCount int64 `json:"total_count"`
	Offset     int   `json:"offset"`
	Limit      int   `json:"limit"`
}

type BookCreateRequest struct {
	Title           string  `json:"title" validate:"required"`
	Description     string  `json:"description"`
	Category        string  `json:"category" validate:"required"`
	Image           string  `json:"image"`
	PublicationYear int     `json:"publication_year"`
	Rating          float64 `json:"rating"`
	Pages           int     `json:"pages"`
	Isbn            string  `json:"isbn"`
	AuthorName      string  `json:"author_name" validate:"required"`
}

type BookUpdateRequest struct {
	Title           *string  `json:"title,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Category        *string  `json:"category,omitempty"`
	Image           *string  `json:"image,omitempty"`
	PublicationYear *int     `json:"publication_year,omitempty"`
	Rating          *float64 `json:"rating,omitempty"`
	Pages           *int     `json:"pages,omitempty"`
	AuthorName      *string  `json:"author_name,omitempty"`
	Isbn            *string  `json:"isbn,omitempty"`
}

type BookResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	Image           string    `json:"image"`
	PublicationYear int       `json:"publication_year"`
	Rating          float64   `json:"rating"`
	Pages           int       `json:"pages"`
	Isbn            string    `json:"isbn"`
	AuthorName      string    `json:"author_name"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

func ToBookResponse(book *models.Book) *BookResponse {
	return &BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}
}

func ToBookListResponse(books []models.Book, meta PaginationMeta) fiber.Map {
	bookResponses := make([]BookResponse, 0)
	for _, book := range books {
		bookResponses = append(bookResponses, *ToBookResponse(&book))
	}

	return fiber.Map{
		"meta": meta,
		"data": bookResponses,
	}
}
