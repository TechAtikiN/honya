package utils

import (
	"regexp"

	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
)

var AllowedOperations = map[string]struct{}{
	"redirection": {},
	"canonical":   {},
	"all":         {},
}

func IsValidUrl(url string) bool {
	const urlRegex = `(?i)^(https)://[^\s/$.?#].[^\s]*$`
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(url)
}

func ValidateProcessUrlRequest(request *dto.ProcessUrlRequest) error {
	if !IsValidUrl(request.Url) {
		return errors.NewBadRequestError("URL is required.")
	}

	if _, valid := AllowedOperations[request.Operation]; !valid {
		return errors.NewBadRequestError("Invalid operation type. Allowed operations are: redirection, canonical, all.")
	}
	return nil
}
