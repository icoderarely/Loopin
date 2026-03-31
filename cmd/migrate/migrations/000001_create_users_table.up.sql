CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    username varchar(50) UNIQUE NOT NULL,
    password BYTEA NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

