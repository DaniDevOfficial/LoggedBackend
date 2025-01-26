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
    );
