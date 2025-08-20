package dtos

type BookCreateRequest struct {
	Title           string  `json:"title" validate:"required"`
	Description     string  `json:"description"`
	Category        string  `json:"category" validate:"required"`
	Image           string  `json:"image"`
	PublicationYear int     `json:"publication_year"`
	Rating          float64 `json:"rating"`
	Pages           int     `json:"pages"`
	Isbn            string  `json:"isbn"`
	AuthorId        string  `json:"author_id" validate:"required"`
}

type BookUpdateRequest struct {
	Title           *string  `json:"title,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Category        *string  `json:"category,omitempty"`
	Image           *string  `json:"image,omitempty"`
	PublicationYear *int     `json:"publication_year,omitempty"`
	Rating          *float64 `json:"rating,omitempty"`
	Pages           *int     `json:"pages,omitempty"`
	Isbn            *string  `json:"isbn,omitempty"`
	AuthorId        *string  `json:"author_id,omitempty"`
}

type BookResponse struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Image           string  `json:"image"`
	PublicationYear int     `json:"publication_year"`
	Rating          float64 `json:"rating"`
	Pages           int     `json:"pages"`
	Isbn            string  `json:"isbn"`
	AuthorId        string  `json:"author_id"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}
