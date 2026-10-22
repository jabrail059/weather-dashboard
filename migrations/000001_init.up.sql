CREATE SCHEMA dashboard;

CREATE TABLE IF NOT EXISTS dashboard.cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE CHECK(char_length(name) BETWEEN 1 AND 30),
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);

CREATE TABLE dashboard.favorites (
    id SERIAL PRIMARY KEY,
    session_id TEXT NOT NULL,
    name VARCHAR(30) NOT NULL CHECK(char_length(name) BETWEEN 1 AND 30),
    created_at TIMESTAMPTZ NOT NULL,
    last_seen_at TIMESTAMPTZ NOT NULL
);
