CREATE TABLE IF NOT EXISTS pager_customer (
    pager_id TEXT,
    name TEXT,
    email TEXT,
    phone TEXT,
    FOREIGN KEY (pager_id) REFERENCES pager(id)
);