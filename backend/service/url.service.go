package service

import (
	"net/url"
	"strings"

	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/utils"
)

type UrlService interface {
	GetRedirectionUrl(url string) (string, error)
	GetCanonicalUrl(url string) (string, error)
}

type urlService struct {
	originalDomain string
}

func NewUrlService(originalDomain string) UrlService {
	return &urlService{
		originalDomain: strings.ToLower(originalDomain),
	}
}

func (s *urlService) GetRedirectionUrl(rawUrl string) (string, error) {
	if !utils.IsValidUrl(rawUrl) {
		return "", errors.NewBadRequestError("Invalid URL format")
	}

	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", errors.NewBadRequestError("Failed to parse URL")
	}

	// Lowercase scheme + host + path
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Host = s.originalDomain
	parsed.Path = strings.ToLower(parsed.Path)

	return parsed.String(), nil
}

func (s *urlService) GetCanonicalUrl(rawUrl string) (string, error) {
	if !utils.IsValidUrl(rawUrl) {
		return "", errors.NewBadRequestError("Invalid URL format")
	}

	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", errors.NewBadRequestError("Failed to parse URL")
	}

	// Remove query params and trailing slashes from path
	parsed.RawQuery = ""
	parsed.Path = strings.TrimRight(parsed.Path, "/")

	return parsed.String(), nil
}
