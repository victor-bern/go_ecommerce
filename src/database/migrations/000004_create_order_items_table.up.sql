CREATE TABLE IF NOT EXISTS order_items(
    id serial PRIMARY KEY,
    order_id serial NOT NULL,
    product_id serial NOT NULL,
    quantity int NOT NULL,
    FOREIGN KEY (order_id) REFERENCES "order"(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);