-- +goose Up
CREATE TABLE user_documents (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    doc_type VARCHAR(100),
    doc_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS user_documents;
