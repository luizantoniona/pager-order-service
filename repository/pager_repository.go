package repository

import (
	"pager-order-service/database"
	"pager-order-service/model"
)

func GetPagerByID(id string) (model.Pager, error) {
	var pager model.Pager

	query, _ := database.ReadSQLFile("./repository/sql/pager/select_pager_by_id.sql")
	err := database.DB.QueryRow(query, id).Scan(&pager.ID, &pager.PagerNumber, &pager.CreatedAt, &pager.UpdatedAt)
	if err != nil {
		return pager, err
	}

	query, _ = database.ReadSQLFile("./repository/sql/pager_customer/select_pager_customer_by_id.sql")
	err = database.DB.QueryRow(query, id).Scan(
		&pager.Customer.Name,
		&pager.Customer.Email,
		&pager.Customer.Phone,
	)
	if err != nil {
		return pager, err
	}

	query, _ = database.ReadSQLFile("./repository/sql/pager_item/select_pager_items_by_id.sql")
	rows, err := database.DB.Query(query, id)
	if err != nil {
		return pager, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.Code,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Quantity,
		)
		if err != nil {
			return pager, err
		}
		pager.Items = append(pager.Items, item)
	}

	return pager, nil
}

func GetAllPagerIDs() ([]string, error) {
	query, _ := database.ReadSQLFile("./repository/sql/pager/select_all_pager_ids.sql")

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func InsertPager(pager model.Pager) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	query, _ := database.ReadSQLFile("./repository/sql/pager/insert_pager.sql")
	_, err = tx.Exec(query, pager.ID, pager.PagerNumber, pager.CreatedAt, pager.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	query, _ = database.ReadSQLFile("./repository/sql/pager_customer/insert_pager_customer.sql")
	_, err = tx.Exec(query, pager.ID, pager.Customer.Name, pager.Customer.Email, pager.Customer.Phone)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range pager.Items {
		query, _ = database.ReadSQLFile("./repository/sql/pager_item/insert_pager_item.sql")
		_, err = tx.Exec(query, pager.ID, item.Code, item.Name, item.Description, item.Price, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
