# Муравьев В.Р. ЭФМО-02-24
## Таблицы интернет-магазина электронники
### Таблица "manufacturers"
```sql
CREATE TABLE manufacturers (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT
);
```
### Таблица "categories"
```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(16)
);
```
### Таблица "contactinfo"
```sql
CREATE TABLE contactinfo (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(11),
    email TEXT,
    first_name TEXT,
    last_name TEXT,
    middle_name TEXT
);
```
### Таблица "products"
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    code VARCHAR(15),
    title TEXT,
    description TEXT,
    price NUMERIC,
    preview VARCHAR(255),
    manufacturer_id INTEGER REFERENCES manufacturers(id),
    category_id INTEGER REFERENCES categories(id)
);
```
### Таблица "users"
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50),
    password TEXT,
    email TEXT,
    phone_number VARCHAR(11),
    is_active BOOLEAN
);
```
### Таблица "carts"
```sql
CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    user_id INTEGER REFERENCES users(id)
);
```
### Таблица "cartproducts"
```sql
CREATE TABLE cartproducts (
    id SERIAL PRIMARY KEY,
    count INTEGER,
    unit_price NUMERIC,
    product_id INTEGER REFERENCES products(id),
    cart_id INTEGER REFERENCES carts(id)
);
```
### Таблица "reviews"
```sql
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    rating SMALLINT,
    user_id INTEGER REFERENCES users(id),
    product_id INTEGER REFERENCES products(id)
);
```
### Таблица "userfavorites"
```sql
CREATE TABLE userfavorites (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    product_id INTEGER REFERENCES products(id)
);
```
### Таблица "orders"
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    code VARCHAR(16),
    status VARCHAR(10),
    created_at TIMESTAMP,
    comment TEXT,
    user_id INTEGER REFERENCES users(id),
    contact_info_id INTEGER REFERENCES contactinfo(id)
);
```
### Таблица "payments"
```sql
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    bank_details TEXT,
    type VARCHAR(10),
    status VARCHAR(10),
    order_id INTEGER REFERENCES orders(id)
);
```
### Таблица "shops"
```sql
CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    address TEXT,
    phone VARCHAR(11),
    opening_hours TEXT,
    status VARCHAR(10)
);
```
### Таблица "remainingstocks"
```sql
CREATE TABLE remainingstocks (
    id SERIAL PRIMARY KEY,
    count INTEGER,
    shop_id INTEGER REFERENCES shops(id),
    product_id INTEGER REFERENCES products(id)
);
```
### Заполнение таблиц данными
```sql
INSERT INTO manufacturers (title, description) VALUES
('Sony', 'A leading manufacturer of electronics and entertainment equipment'),
('Samsung', 'South Korean multinational electronics company'),
('Apple', 'American multinational technology company that designs and sells electronics');

INSERT INTO categories (title) VALUES
('Smartphones'),
('Laptops'),
('Accessories');

INSERT INTO contactinfo (phone_number, email, first_name, last_name, middle_name) VALUES
('89991234567', 'john.doe@example.com', 'John', 'Doe', 'Edwardovich'),
('89992345678', 'jane.smith@example.com', 'Jane', 'Smith', NULL);

INSERT INTO products (code, title, description, price, preview, manufacturer_id, category_id) VALUES
('P001', 'iPhone 14', 'Latest model of iPhone', 99999.99, 'iphone14.jpg', 3, 1),
('P002', 'Galaxy S22', 'Samsung flagship smartphone', 79999.99, 'galaxy_s22.jpg', 2, 1),
('P003', 'MacBook Air', 'Lightweight laptop from Apple', 119999.99, 'macbook_air.jpg', 3, 2);

INSERT INTO users (login, password, email, phone_number, is_active) VALUES
('john_doe', 'password123', 'john.doe@example.com', '89991234567', TRUE),
('jane_smith', 'securepass', 'jane.smith@example.com', '89992345678', TRUE);

INSERT INTO carts (created_at, user_id) VALUES
(NOW(), 1),
(NOW(), 2);

INSERT INTO cartproducts (count, unit_price, product_id, cart_id) VALUES
(1, 99999.99, 1, 1),
(2, 79999.99, 2, 2);

INSERT INTO reviews (title, description, rating, user_id, product_id) VALUES
('Great phone', 'I love the new features!', 5, 1, 1),
('Worth the price', 'Excellent performance', 4, 2, 2);

INSERT INTO userfavorites (user_id, product_id) VALUES
(1, 1),
(2, 2);

INSERT INTO orders (code, status, created_at, comment, user_id, contact_info_id) VALUES
('ORD001', 'Pending', NOW(), '', 1, 1),
('ORD002', 'Completed', NOW(), '', 2, 2);

INSERT INTO orderproducts (count, unit_price, order_id, product_id) VALUES
(1, 99999.99, 1, 1),
(1, 119999.99, 2, 3);

INSERT INTO payments (bank_details, type, status, order_id) VALUES
('56781234', 'Credit', 'Paid', 1),
('56785678', 'Credit', 'Paid', 2);

INSERT INTO shops (address, phone, opening_hours, status) VALUES
('123 Main St, Moscow', '89991112233', '9:00-21:00', 'Open'),
('456 Elm St, St. Petersburg', '89992223344', '10:00-22:00', 'Open');

INSERT INTO remainingstocks (count, shop_id, product_id) VALUES
(10, 1, 1),
(5, 1, 2),
(7, 2, 3);
```
