CREATE TABLE IF NOT EXISTS product_inventory(
    id serial PRIMARY KEY,
    product_id serial,
    stock int NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)