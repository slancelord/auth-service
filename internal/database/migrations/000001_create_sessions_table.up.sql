DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    guid UUID NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL
);