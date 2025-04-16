SELECT 
    id,
    pager_number,
    created_at,
    updated_at
FROM
    pager
WHERE
    id = ?;
