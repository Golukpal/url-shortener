package models

import "time"

type URL struct {
	ID          string     `db:"id" json:"id"`
	OriginalURL string     `db:"original_url" json:"original_url"`
	ShortCode   string     `db:"short_code" json:"short_code"`
	Clicks      int        `db:"clicks" json:"clicks"`
	ExpiresAt   *time.Time `db:"expires_at" json:"expires_at,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}
