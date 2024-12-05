-- create tables

CREATE TABLE game_boxes (
    game_boxes_id BIGSERIAL PRIMARY KEY,
    amount NUMERIC(14,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE source_of_funds (
    source_of_fund_id BIGSERIAL PRIMARY KEY,
    source_name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE transaction_types (
    transaction_type_id BIGSERIAL PRIMARY KEY,
    type_name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR NOT NULL,
    user_password VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    profile_image VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE wallets (
    wallet_id BIGSERIAL PRIMARY KEY,
    wallet_number VARCHAR NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(user_id),
    amount NUMERIC(14,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE game_attempts (
    game_attempt_id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT NOT NULL REFERENCES wallets(wallet_id),
    amount NUMERIC(14,2) NOT NULL,
    game_boxes_id BIGINT NOT NULL REFERENCES game_boxes(game_boxes_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE transactions (
    transaction_id BIGSERIAL PRIMARY KEY,
    transaction_ref_id BIGINT REFERENCES transactions(transaction_id),
    wallet_id BIGINT NOT NULL REFERENCES wallets(wallet_id),
    transaction_additional_detail_id BIGINT NOT NULL,
    transaction_type_id BIGINT NOT NULL REFERENCES transaction_types(transaction_type_id),
    amount NUMERIC(14,2) NOT NULL,
    transaction_description VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE reset_password_attempts (
    reset_password_attempt_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id),
    reset_code VARCHAR NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_unique_email_not_deleted 
ON users(email) 
WHERE deleted_at IS NULL;

INSERT INTO public.game_boxes (amount, created_at, updated_at) VALUES
	 (50000.00, NOW(), NOW()),
	 (500000.00, NOW(), NOW()),
	 (10000.00, NOW(), NOW()),
	 (100000.00, NOW(), NOW()),
	 (1000000.00, NOW(), NOW()),
	 (200000.00, NOW(), NOW()),
	 (20000.00, NOW(), NOW()),
	 (2000000.00, NOW(), NOW()),
	 (5000.00, NOW(), NOW());

INSERT INTO public.source_of_funds (source_name, created_at, updated_at) VALUES
	 ('BANK_TRANSFER', NOW(), NOW()),
	 ('CREDIT_CARD', NOW(), NOW()),
	 ('CASH', NOW(), NOW()),
	 ('REWARD', NOW(), NOW());

INSERT INTO public.transaction_types (type_name, created_at, updated_at) VALUES
	 ('TOP_UP', NOW(), NOW()),
	 ('TRANSFER', NOW(), NOW());