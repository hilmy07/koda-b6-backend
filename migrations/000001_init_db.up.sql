CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(80) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    fullname VARCHAR(50),
    phone VARCHAR(30),
    address VARCHAR(100),
    profile_img VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name_product VARCHAR(150) NOT NULL,
    description TEXT,
    base_price DECIMAL(12,2) NOT NULL,
    stock INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_variants (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    variant_name VARCHAR(100),
    add_price DECIMAL(12,2) DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_sizes (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    size_name VARCHAR(100),
    add_price DECIMAL(12,2) DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    path VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_discounts (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    is_flashsale BOOLEAN DEFAULT FALSE,
    discount_rate DECIMAL(5,2),
    description TEXT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_reviews (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    message TEXT,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users_coffee(id) ON DELETE CASCADE,
    transaction_code VARCHAR(100),
    delivery_method VARCHAR(100),
    fullname VARCHAR(100),
    email VARCHAR(100),
    address TEXT,
    subtotal DECIMAL(12,2),
    delivery_fee DECIMAL(12,2),
    tax DECIMAL(12,2),
    total DECIMAL(12,2),
    status VARCHAR(50),
    payment_method VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_items (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    variant_id INT REFERENCES variants(id),
    size_id INT REFERENCES sizes(id),
    quantity INT,
    price DECIMAL(12,2)
);

SELECT * FROM sizes;
SELECT * FROM variants;
SELECT * FROM users;
SELECT * FROM reviews;
SELECT * FROM product_images;



