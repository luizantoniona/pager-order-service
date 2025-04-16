package repository

import (
	"pager-order-service/database"
	"pager-order-service/model"
)

func InsertOrder(order model.Order) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	query, _ := database.ReadSQLFile("./repository/sql/order/insert_order.sql")
	_, err = tx.Exec(query, order.ID, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	query, _ = database.ReadSQLFile("./repository/sql/order_customer/insert_order_customer.sql")
	_, err = tx.Exec(query, order.ID, order.Customer.Name, order.Customer.Email, order.Customer.Phone)
	if err != nil {
		tx.Rollback()
		return err
	}

	query, _ = database.ReadSQLFile("./repository/sql/order_address/insert_order_address.sql")
	_, err = tx.Exec(query, order.ID, order.Address.Street, order.Address.City, order.Address.State,
		order.Address.ZipCode, order.Address.Country, order.Address.Observations)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		query, _ = database.ReadSQLFile("./repository/sql/order_item/insert_order_item.sql")
		_, err = tx.Exec(query, order.ID, item.Code, item.Name, item.Description, item.Price, item.Quantity)
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
