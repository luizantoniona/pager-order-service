SELECT
    street,
    city,
    state,
    zip_code,
    country,
    observations
FROM
    order_address
WHERE
    order_id = ?;
