CREATE TABLE IF NOT EXISTS order_address (
    order_id TEXT,
    street TEXT,
    city TEXT,
    state TEXT,
    zip_code TEXT,
    country TEXT,
    observations TEXT,
    FOREIGN KEY (order_id) REFERENCES "order"(id)
);
