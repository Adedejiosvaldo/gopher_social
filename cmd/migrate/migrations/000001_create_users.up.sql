CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email citext NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password bytea NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
);
