package service_test

import (
	"honya/backend/errors"
	"honya/backend/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlService_GetRedirectionUrl(t *testing.T) {
	originalDomain := "www.example.com"
	svc := service.NewUrlService(originalDomain)

	tests := []struct {
		name      string
		inputUrl  string
		expected  string
		expectErr bool
	}{
		{
			name:      "valid http url",
			inputUrl:  "https://OldDomain.com/Path/To/Page",
			expected:  "https://www.example.com/path/to/page",
			expectErr: false,
		},
		{
			name:      "valid url with mixed case",
			inputUrl:  "HTTPS://SomeDomain.com/TEST",
			expected:  "https://www.example.com/test",
			expectErr: false,
		},
		{
			name:      "invalid url",
			inputUrl:  "htp://invalid-url",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "empty url",
			inputUrl:  "",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.GetRedirectionUrl(tt.inputUrl)
			if tt.expectErr {
				assert.Error(t, err)
				assert.IsType(t, &errors.AppError{}, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestUrlService_GetCanonicalUrl(t *testing.T) {
	svc := service.NewUrlService("www.example.com")

	tests := []struct {
		name      string
		inputUrl  string
		expected  string
		expectErr bool
	}{
		{
			name:      "url with query params",
			inputUrl:  "https://example.com/path/to/page/?q=test",
			expected:  "https://example.com/path/to/page",
			expectErr: false,
		},
		{
			name:      "url with trailing slash",
			inputUrl:  "https://example.com/path/",
			expected:  "https://example.com/path",
			expectErr: false,
		},
		{
			name:      "invalid url",
			inputUrl:  "htp://invalid-url",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "empty url",
			inputUrl:  "",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.GetCanonicalUrl(tt.inputUrl)
			if tt.expectErr {
				assert.Error(t, err)
				assert.IsType(t, &errors.AppError{}, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
