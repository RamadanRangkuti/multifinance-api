-- +goose Up
CREATE TABLE user_limits (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    tenor INT NOT NULL,
    limit_amount NUMERIC(10,2) NOT NULL,
    used_amount NUMERIC(10,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS user_limits;
