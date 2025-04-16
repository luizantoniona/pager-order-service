SELECT 
    name,
    email,
    phone
FROM
    order_customer
WHERE
    order_id = ?;
