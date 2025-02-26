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
    message TEXT,
    request TEXT,
    user_id TEXT, -- Just an identifier, no FK
    request_url TEXT,
    response TEXT,
    life_time TIMESTAMPTZ,
    request_key TEXT,
    date_time TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS "refreshTokens" (
                                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    refresh_token VARCHAR(255) NOT NULL,
    life_time TIMESTAMPTZ DEFAULT NULL,
    last_usage TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS "userHasRoles" (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id UUID REFERENCES users(id) ON DELETE CASCADE,
        role VARCHAR(255) REFERENCES roles(role) ON DELETE CASCADE,
    )

CREATE TABLE IF NOT EXISTS roles (
    role VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

-- Insert an admin user (unclaimed) with the password 'admin'
INSERT INTO roles (role)
VALUES ("admin");

WITH inserted_user AS (
    INSERT INTO users (id, username, password, is_claimed)
    VALUES (gen_random_uuid(), 'admin', '$2a$10$AJGqGxn0Cj3wAFsLpl6jjeT.cD3ipzZXAvE5pLychFvDrhygt63mi', false)
    RETURNING id
)

INSERT INTO userHasRole (user_id, role)
SELECT id, 'admin' FROM inserted_user;
