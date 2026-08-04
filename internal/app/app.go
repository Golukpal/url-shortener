package app

import (
	"github.com/Golukpal/url-shortener/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *pgxpool.Pool
}

func New(
	cfg *config.Config,
	logger *zap.Logger,
	db *pgxpool.Pool,
) *App {
	return &App{
		Config: cfg,
		Logger: logger,
		DB:     db,
	}
}
