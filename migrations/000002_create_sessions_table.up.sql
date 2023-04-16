CREATE TABLE IF NOT EXISTS sessions
(
    "refresh_token" uuid PRIMARY KEY,
    "user_id"       uuid REFERENCES users (id) ON DELETE CASCADE,
    "expires_in"    timestamp with time zone NOT NULL
);