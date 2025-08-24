package dto

type ProcessUrlRequest struct {
	Url       string `json:"url" validate:"required,url"`
	Operation string `json:"operation" validate:"required,oneof=redirection canonical all"`
}

type ProcessUrlResponse struct {
	ProcessedUrl string `json:"processed_url"`
}
