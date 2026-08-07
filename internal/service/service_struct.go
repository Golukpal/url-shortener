package service

import (
	"context"

	"github.com/Golukpal/url-shortener/internal/models"
	"github.com/Golukpal/url-shortener/internal/repository"
)

type urlService struct {
	repo repository.URLRepository
}

func (s *urlService) Delete(ctx context.Context, shortCode string) error {
	return s.repo.Delete(ctx, shortCode)
}

func (s *urlService) GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	return s.repo.GetByShortCode(ctx, shortCode)
}

func (s *urlService) Redirect(ctx context.Context, shortCode string) (*models.URL, error) {
	url, err := s.repo.GetByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {

		return nil, err

	}

	err = s.repo.IncreamentClicks(
		ctx,
		shortCode,
	)

	if err != nil {

		return nil, err

	}

	return url, nil
}

func NewURLService(
	repo repository.URLRepository,
) URLService {

	return &urlService{
		repo: repo,
	}
}

func (s *urlService) Create(

	ctx context.Context,

	originalURL string,

) (*models.URL, error) {
	if err := ValidateURL(originalURL); err != nil {
		return nil, err
	}

	url := &models.URL{

		OriginalURL: originalURL,

		ShortCode: GenerateShortCode(),
	}

	if err := s.repo.Create(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}
