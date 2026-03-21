CREATE TABLE if NOT EXISTS product_variants (
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    variant_id INT NOT NULL REFERENCES variants(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, variant_id)
);