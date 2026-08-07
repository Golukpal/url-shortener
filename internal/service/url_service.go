package service

import (
	"context"

	"github.com/Golukpal/url-shortener/internal/models"
)

type URLService interface {
	Create(
		ctx context.Context,
		originalUrl string,
	) (*models.URL, error)

	GetByShortCode(
		ctx context.Context,
		shortCode string,
	)(*models.URL, error)

	Delete(
		ctx context.Context,
		shortCode string,
	) error

	Redirect(
		ctx context.Context,
		shortCode string,
	)(*models.URL, error)
}
