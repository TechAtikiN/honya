package utils

import (
	"regexp"

	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
)

var allowedOperations = map[string]struct{}{
	"redirection": {},
	"canonical":   {},
	"all":         {},
}

func IsValidUrl(url string) bool {
	const urlRegex = `^(https)://[^\s/$.?#].[^\s]*$`
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(url)
}

func ValidateProcessUrlRequest(request *dto.ProcessUrlRequest) error {
	if !IsValidUrl(request.Url) {
		return errors.NewBadRequestError("URL is required.")
	}

	if _, valid := allowedOperations[request.Operation]; !valid {
		return errors.NewBadRequestError("Invalid operation type. Allowed operations are: redirection, canonical, all.")
	}
	return nil
}
