SELECT
    code,
    name,
    description,
    price,
    quantity
FROM
    order_item
WHERE
    order_id = ?;
