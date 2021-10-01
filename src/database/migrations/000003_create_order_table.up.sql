CREATE TABLE IF NOT EXISTS "order" (
    id serial PRIMARY KEY,
    user_id serial,
    total FLOAT(6) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);