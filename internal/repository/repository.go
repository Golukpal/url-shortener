package repository

import (
	"context"

	"github.com/Golukpal/url-shortener/internal/models"
)

type URLRepository interface {
	Create(
		ctx context.Context,
		url *models.URL,
	)error

	GetByShortCode(
		ctx context.Context,
		shortCode string, 
	)(*models.URL, error)

	Delete(
		ctx context.Context,
		shortCode string,

	)error

	IncreamentClicks(
		ctx context.Context,
		shortCode string,
	)error
}