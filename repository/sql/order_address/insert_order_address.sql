INSERT INTO order_address (
    order_id,
    street,
    city,
    state,
    zip_code,
    country,
    observations
)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);
