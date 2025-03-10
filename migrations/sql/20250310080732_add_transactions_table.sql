-- +goose Up
CREATE TABLE transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    tenor INT NOT NULL,
    contract_number VARCHAR(50) NOT NULL UNIQUE,
    otr NUMERIC(10,2) NOT NULL,
    admin_fee NUMERIC(15,2) NOT NULL,
    installment_count INT NOT NULL,
    interest NUMERIC(5,2) NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    asset_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS transactions;
