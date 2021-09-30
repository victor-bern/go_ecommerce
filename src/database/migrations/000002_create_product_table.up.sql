CREATE TABLE IF NOT EXISTS product (
   id serial PRIMARY KEY,
   title VARCHAR (250) NOT NULL,
   description VARCHAR (250) NOT NULL,
   price FLOAT (6) NOT NULL,
   created_at TIMESTAMPTZ DEFAULT NOW(),
   updated_at TIMESTAMPTZ DEFAULT NOW()
);