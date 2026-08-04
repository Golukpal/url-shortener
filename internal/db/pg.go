package db

import (
	"context"
	"time"

	"github.com/Golukpal/url-shortener/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return pgxpool.New(
		ctx,
		cfg.DatabaseURL(),
	)
}
