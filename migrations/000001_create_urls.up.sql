CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS urls (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    original_url TEXT NOT NULL,

    short_code VARCHAR(10) UNIQUE NOT NULL,

    clicks INTEGER NOT NULL DEFAULT 0,

    expires_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()

);