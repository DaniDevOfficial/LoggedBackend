-- CREATE DATABASE logged_in;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_claimed BOOLEAN DEFAULT false,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS logs (
                                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    severity VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    request TEXT NOT NULL,
    user_id UUID NOT NULL, -- Just an identifier, no FK
    request_url TEXT NOT NULL,
    response TEXT NOT NULL,
    life_time INTERVAL NOT NULL,
    request_key TEXT NOT NULL,
    date_time TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS refreshTokens (
                                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    refresh_token VARCHAR(255) NOT NULL,
    life_time INTERVAL NOT NULL,
    last_usage TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

-- Insert an admin user (unclaimed) with the password 'admin'
INSERT INTO users (id, username, password, is_claimed)
VALUES (gen_random_uuid(), 'admin', '$2a$10$AJGqGxn0Cj3wAFsLpl6jjeT.cD3ipzZXAvE5pLychFvDrhygt63mi', false);