package utils

import "time"

const (
	DefaultOffset          = 0
	DefaultLimit           = 10
	DefaultPublicationYear = 2025
	DefaultRating          = 0.0
	DefaultPages           = 0
)

const (
	RateLimitMaxRequests    = 100
	RateLimitExpiryDuration = 1 * time.Minute
)

const (
	OpRedirection = "redirection"
	OpCanonical   = "canonical"
	OpAll         = "all"
)

const (
	DefaultDonutChartFilterBy = "category"
)
