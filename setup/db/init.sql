CREATE DATABASE IF NOT EXISTS logged_in;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS logs (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    severity VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    request TEXT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    request_url TEXT NOT NULL,
    response TEXT NOT NULL,
    life_time VARCHAR(50) NOT NULL,
    request_key VARCHAR(255) NOT NULL,
    date_time TIMESTAMP NOT NULL
    created_at TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS refreshTokens (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id REFERENCES users (user_id) ON DELETE SET NULL
    refreshToken VARCHAR(255) NOT NULL
    life_time TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
    last_usage TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
    created_at TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS users (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL
    password VARCHAR(255) NOT NULL
    is_claimed BOOLEAN DEFAULT False

    last_updated TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
    created_at TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
)


