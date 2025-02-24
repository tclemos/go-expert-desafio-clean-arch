CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    customer_name TEXT,
    status TEXT,
    amount INTEGER,
    items TEXT
);