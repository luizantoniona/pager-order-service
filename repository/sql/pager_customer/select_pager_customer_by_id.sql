SELECT 
    name,
    email,
    phone
FROM
    pager_customer
WHERE
    pager_id = ?;
