package repository

import (
	"pager-order-service/database"
	"pager-order-service/model"
)

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
