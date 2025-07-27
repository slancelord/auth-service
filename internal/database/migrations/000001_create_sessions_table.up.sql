CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    refresh_token TEXT NOT NULL,
    token_pair_id UUID NOT NULL,
    user_agent TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL
);