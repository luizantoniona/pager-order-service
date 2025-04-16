SELECT
    code,
    name,
    description,
    price,
    quantity
FROM
    pager_item
WHERE
    pager_id = ?;
