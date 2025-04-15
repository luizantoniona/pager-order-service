CREATE TABLE IF NOT EXISTS pager_item (
    pager_id TEXT,
    code TEXT,
    name TEXT,
    description TEXT,
    price REAL,
    quantity INTEGER,
    FOREIGN KEY (pager_id) REFERENCES pager(id)
);