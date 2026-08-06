package repository

import (
	"context"

	"github.com/Golukpal/url-shortener/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	db *pgxpool.Pool
}

func NewPGRepository(db *pgxpool.Pool) URLRepository {
	return &PGRepo{
		db: db,
	}
}


func (r *PGRepo) Delete(ctx context.Context, shortCode string) error {
	query:= `DELETE FROM urls WHERE short_code = &1`
	_, err := r.db.Exec(
		ctx, query, shortCode,)
	return err
}


func (r *PGRepo) GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	query := `
	SELECT
		id,
		original_url,
		short_code,
		clicks,
		expires_at,
		created_at,
		updated_at
	FROM urls
	WHERE short_code = $1
	`

	var url models.URL

	err := r.db.QueryRow(
		ctx,
		query,
		shortCode,
	).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.Clicks,
		&url.ExpiresAt,
		&url.CreatedAt,
		&url.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &url, nil

}


func (r *PGRepo) IncreamentClicks(ctx context.Context, shortCode string) error {
	query := `
	UPDATE urls
	SET
		clicks = clicks + 1,
		updated_at = NOW()
	WHERE short_code = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		shortCode,
	)

	return err
}



func (r *PGRepo) Create(ctx context.Context,url *models.URL,) error {
	query := `
		INSERT INTO  urls (original_url, short_code, expires_at)
		VALUES ($1, $2, $3)`
	_, err:= r.db.Exec(
		ctx,
		query,
		url.OriginalURL,
		url.ShortCode,
		url.ExpiresAt,
	)

	return err
}
