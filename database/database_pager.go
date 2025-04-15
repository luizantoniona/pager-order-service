package database

import (
	"log"
)

func createPagerTable() {
	query := `
	CREATE TABLE IF NOT EXISTS pager (
		id TEXT PRIMARY KEY,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createPagerCustomerTable() {
	query := `
	CREATE TABLE IF NOT EXISTS pager_customer (
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

func createPagerItemTable() {
	query := `
	CREATE TABLE IF NOT EXISTS pager_item (
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
