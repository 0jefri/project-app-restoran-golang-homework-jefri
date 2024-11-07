CREATE DATABASE resto_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    role VARCHAR(10) NOT NULL, -- 'Admin', 'Koki', 'Pelanggan'
    password VARCHAR(50) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    order_status VARCHAR(20) NOT NULL, -- 'sedang diproses', 'selesai', 'dibatalkan'
    total_price DECIMAL(10, 2) NOT NULL,
    discount_code VARCHAR(20),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (username, role, password) VALUES
('admin_user', 'Admin', 'password123'),
('koki_user1', 'Koki', 'password123'),
('koki_user2', 'Koki', 'password123'),
('pelanggan_user1', 'Pelanggan', 'password123'),
('pelanggan_user2', 'Pelanggan', 'password123');

INSERT INTO orders (user_id, order_status, total_price, discount_code, rating) VALUES
(1, 'sedang diproses', 100.00, NULL, 5),
(2, 'selesai', 200.00, 'DISC10', 4),
(3, 'dibatalkan', 150.00, NULL, NULL),
(4, 'sedang diproses', 120.00, 'DISC15', 3),
(5, 'selesai', 180.00, NULL, 4);
