CREATE TABLE IF NOT EXISTS cart (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    product_id INT REFERENCES products(id),
    size_id INT REFERENCES sizes(id),
    variant_id INT REFERENCES variants(id),
    quantity INT NOT NULL CHECK(quantity > 0)
);