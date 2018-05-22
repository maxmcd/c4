CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    phone INT,
    country_code INT,
    username TEXT UNIQUE,
    nexmo_request_id TEXT,
    games_played INT DEFAULT 0,
    rating_scratch INT DEFAULT 0
);

CREATE UNIQUE INDEX ON users (phone, country_code);
