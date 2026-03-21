CREATE TABLE if NOT EXISTS product_sizes (
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    size_id INT NOT NULL REFERENCES sizes(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, size_id)
);