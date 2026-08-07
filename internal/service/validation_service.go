package service

import (
	"net/url"

	"github.com/Golukpal/url-shortener/internal/errors"
)

func ValidateURL(rawURL string) error {

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.ErrInvalidURL
	}

	if parsedURL.Scheme == "" {
		return errors.ErrInvalidURL
	}

	if parsedURL.Host == "" {
		return errors.ErrInvalidURL
	}

	return nil
}
