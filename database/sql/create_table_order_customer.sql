CREATE TABLE IF NOT EXISTS order_customer (
    order_id TEXT,
    name TEXT,
    email TEXT,
    phone TEXT,
    FOREIGN KEY (order_id) REFERENCES "order"(id)
);
