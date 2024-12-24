DROP DATABASE IF EXISTS mytheresa;

CREATE DATABASE mytheresa;

\connect mytheresa;

-- create table products
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(6) NOT NULL,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(20) NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted TIMESTAMP
);

-- add indices for category and price
CREATE INDEX category_idx ON products (category);
CREATE INDEX price_idx ON products (price);

-- possible improvements:
-- add unique index for sku field,
-- create new table for categories (id, name), and alter products.category to category_id

-- insert initial products
INSERT INTO products (sku, name, category, price) VALUES 
('000001','BV Lean leather ankle boots', 'boots', 89000),
('000002','BV Lean leather ankle boots', 'boots', 99000),
('000003','Ashlington leather ankle boots', 'boots', 71000),
('000004','Naima embellished suede sandals', 'sandals', 79500),
('000005','Nathane leather sneakers', 'sneakers', 59000),
('000006','Tom Ford Suede blazer', 'blazers', 159000)
;
