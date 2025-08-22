package dto

type QueryParams struct {
	Query  string
	Limit  int
	Offset int
}

type PaginationMeta struct {
	TotalCount int64 `json:"total_count"`
	Offset     int   `json:"offset"`
	Limit      int   `json:"limit"`
}
