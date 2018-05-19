CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    phone INT,
    country_code INT,
    username TEXT UNIQUE,
    nexmo_request_id TEXT,
    rating INT DEFAULT 150000,
    INDEX (phone),
    INDEX (username),
    INDEX (nexmo_request_id)
);

CREATE UNIQUE INDEX ON users (phone, country_code);
