CREATE TABLE IF NOT EXISTS order_item (
    order_id TEXT,
    code TEXT,
    name TEXT,
    description TEXT,
    price REAL,
    quantity INTEGER,
    FOREIGN KEY (order_id) REFERENCES "order"(id)
);
