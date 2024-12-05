DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS ResetPasswords CASCADE;
DROP TABLE IF EXISTS Wallets CASCADE;
DROP TABLE IF EXISTS TransactionUsers CASCADE;
DROP TABLE IF EXISTS SourceFunds CASCADE;
DROP TABLE IF EXISTS Transactions CASCADE;
DROP TABLE IF EXISTS game_boxs CASCADE;

CREATE TABLE Users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    chance_game INTEGER DEFAULT 0 NOT NULL,
    fullname VARCHAR NOT NULL,
    birthdate DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE ResetPasswords (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    token VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE Wallets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    wallet_number VARCHAR NOT NULL,
    balance DECIMAL DEFAULT 0 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE TransactionUsers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL ,
    transaction_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE SourceFunds (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE Transactions (
    id BIGSERIAL PRIMARY KEY,
    source_id BIGINT NOT NULL,
    recipient_id BIGINT NOT NULL,
    transaction_time TIMESTAMP NOT NULL,
    amount DECIMAL NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE game_boxs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    is_open BOOLEAN NOT NULL,
    box1 INTEGER NOT NULL,
    box2 INTEGER NOT NULL,
    box3 INTEGER NOT NULL,
    box4 INTEGER NOT NULL,
    box5 INTEGER NOT NULL,
    box6 INTEGER NOT NULL,
    box7 INTEGER NOT NULL,
    box8 INTEGER NOT NULL,
    box9 INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

ALTER TABLE ResetPasswords 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE Wallets 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE TransactionUsers 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE TransactionUsers 
ADD FOREIGN KEY (transaction_id) REFERENCES Transactions(id);

ALTER TABLE Transactions 
ADD FOREIGN KEY (source_id) REFERENCES SourceFunds(id);

ALTER TABLE Transactions 
ADD FOREIGN KEY (recipient_id) REFERENCES Users(id);

ALTER TABLE game_boxs 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

INSERT INTO Users (email, password, chance_game, fullname, birthdate, created_at, updated_at)
VALUES 
('frieren@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 2, 'Paduka Ratu Frieren', '1990-01-01', NOW(), NOW()),
('pogba@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 0, 'Ahmad Pogba', '1995-04-05', NOW(), NOW()),
('ibra@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 0, 'Ibrahimopiko', '1977-11-05', NOW(), NOW()),
('messi@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 0, 'Messi Udin', '1986-02-02', NOW(), NOW()),
('ronaldo@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 1, 'Cristiano Ronaldo', '1992-02-02', NOW(), NOW());

INSERT INTO ResetPasswords (user_id, token, created_at, updated_at)
VALUES 
(1, '202401101606205', NOW(), NOW()),
(2, '202402101606204', NOW(), NOW()),
(3, '202403101606203', NOW(), NOW()),
(4, '202404101606202', NOW(), NOW()),
(5, '202407101606201', NOW(), NOW());

INSERT INTO Wallets (user_id, wallet_number, balance, created_at, updated_at)
VALUES 
(1, '7770000000001', 1000.00, NOW(), NOW()),
(2, '7770000000002', 2000.00, NOW(), NOW()),
(3, '7770000000003', 3000.00, NOW(), NOW()),
(4, '7770000000004', 4000.00, NOW(), NOW()),
(5, '7770000000005', 5000.00, NOW(), NOW());

INSERT INTO SourceFunds (name, created_at, updated_at)
VALUES 
('Bank Transfer', NOW(), NOW()),
('Credit Card', NOW(), NOW()),
('Cash', NOW(), NOW()),
('Reward', NOW(), NOW());

INSERT INTO Transactions (source_id, recipient_id, transaction_time, amount, description, created_at, updated_at)
VALUES 
(1, 2, '2024-01-01', 500.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(2, 1, '2024-02-01', 150.00, 'Top Up from Credit Card', NOW(), NOW()),
(3, 3, '2024-03-02', 300.00, 'Top Up from Cash', NOW(), NOW()),
(1, 3, '2024-03-03', 750.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(3, 3, '2024-04-03', 900.00, 'Top Up from Cash', NOW(), NOW()),
(1, 2, '2024-05-01', 1500.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(1, 1, '2024-06-11', 200.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(3, 3, '2024-06-12', 1000.00, 'Top Up from Cash', NOW(), NOW()),
(4, 3, '2024-07-10', 9000.00, 'Top Up from Reward', NOW(), NOW()),
(1, 3, '2024-08-05', 320.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(2, 2, '2024-08-03', 5000.00, 'Top Up from Credit Card', NOW(), NOW()),
(1, 1, '2024-09-01', 2000.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(3, 3, '2024-09-04', 140.00, 'Top Up from Cash', NOW(), NOW()),
(4, 3, '2024-10-07', 20000.00, 'Top Up from Reward', NOW(), NOW()),
(2, 3, '2024-10-08', 2400.00, 'Top Up from Credit Card', NOW(), NOW()),
(4, 2, '2024-11-10', 4500.00, 'Top Up from Reward', NOW(), NOW()),
(1, 1, '2024-11-11', 550.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(4, 3, '2024-12-01', 9000.00, 'Top Up from Reward', NOW(), NOW()),
(1, 3, '2024-12-05', 250.00, 'Top Up from Bank Transfer', NOW(), NOW()),
(3, 3, '2024-12-07', 2500.00, 'Top Up from Cash', NOW(), NOW());

INSERT INTO TransactionUsers (user_id, transaction_id, created_at, updated_at)
VALUES 
(1, 1, NOW(), NOW()),
(2, 2, NOW(), NOW()),
(1, 3, NOW(), NOW()),
(4, 4, NOW(), NOW()),
(1, 5, NOW(), NOW()),
(1, 6, NOW(), NOW()),
(2, 7, NOW(), NOW()),
(1, 8, NOW(), NOW()),
(1, 9, NOW(), NOW()),
(5, 10, NOW(), NOW()),
(1, 11, NOW(), NOW()),
(3, 12, NOW(), NOW()),
(1, 13, NOW(), NOW()),
(4, 14, NOW(), NOW()),
(1, 15, NOW(), NOW()),
(1, 16, NOW(), NOW()),
(5, 17, NOW(), NOW()),
(3, 18, NOW(), NOW()),
(4, 19, NOW(), NOW()),
(1, 20, NOW(), NOW());


