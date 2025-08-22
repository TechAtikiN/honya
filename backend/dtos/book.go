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
	AuthorName      string  `json:"author_name"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}

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
