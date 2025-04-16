package repository

import (
	"pager-order-service/database"
	"pager-order-service/model"
)

func GetOrderByID(id string) (model.Order, error) {
	var order model.Order

	query, _ := database.ReadSQLFile(".repository/sql/order/select_order_by_id.sql")
	err := database.DB.QueryRow(query, id).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return order, err
	}

	query, _ = database.ReadSQLFile(".repository/sql/order_customer/select_order_customer_by_id.sql")
	err = database.DB.QueryRow(query, id).Scan(
		&order.Customer.Name,
		&order.Customer.Email,
		&order.Customer.Phone,
	)
	if err != nil {
		return order, err
	}

	query, _ = database.ReadSQLFile(".repository/sql/order_address/select_order_address_by_id.sql")
	err = database.DB.QueryRow(query, id).Scan(
		&order.Address.Street,
		&order.Address.City,
		&order.Address.State,
		&order.Address.ZipCode,
		&order.Address.Country,
		&order.Address.Observations,
	)
	if err != nil {
		return order, err
	}

	query, _ = database.ReadSQLFile(".repository/sql/order_item/select_order_items_by_id.sql")
	rows, err := database.DB.Query(query, id)
	if err != nil {
		return order, err
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
			return order, err
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func GetAllOrderIDs() ([]string, error) {

	query, _ := database.ReadSQLFile(".repository/sql/order/select_all_order_ids.sql")
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

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
