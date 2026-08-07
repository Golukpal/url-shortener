package errors

import "errors"

var (
	ErrInvalidURL         = errors.New("invalid URL")
	ErrURLNotFound        = errors.New("URL not found")
	ErrDuplicateShortCode = errors.New("duplicate short code")
)
