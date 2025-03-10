-- +goose Up
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    nik VARCHAR(50) NOT NULL UNIQUE,
    fullname VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    place_of_birth VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    salary NUMERIC(10,2) NOT NULL,
    password VARCHAR(255) NOT NULL, 
    role VARCHAR(50) NOT NULL DEFAULT 'User',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
