package database

import (
	"log"
)

func createOrderTable() {
	query := `
	CREATE TABLE IF NOT EXISTS order (
		id TEXT PRIMARY KEY,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderCustomerTable() {
	query := `
	CREATE TABLE IF NOT EXISTS order_customer (
		order_id TEXT,
		name TEXT,
		email TEXT,
		phone TEXT,
		FOREIGN KEY (order_id) REFERENCES order(id)
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderAddressTable() {
	query := `
	CREATE TABLE IF NOT EXISTS order_address (
		order_id TEXT,
		street TEXT,
		city TEXT,
		state TEXT,
		zip_code TEXT,
		country TEXT,
		observations TEXT,
		FOREIGN KEY (order_id) REFERENCES order(id)
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderItemTable() {
	query := `
	CREATE TABLE IF NOT EXISTS order_item (
		order_id TEXT,
		code TEXT,
		name TEXT,
		description TEXT,
		price REAL,
		quantity INTEGER,
		FOREIGN KEY (order_id) REFERENCES order(id)
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
