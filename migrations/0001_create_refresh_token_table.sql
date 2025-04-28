CREATE TABLE refresh_token (
    user_id UUID PRIMARY KEY,
    token_hash TEXT NOT NULL,
    ip TEXT NOT NULL
);