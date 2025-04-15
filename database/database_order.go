package database

import (
	"log"
	"os"
)

func createOrderTable() {
	query, err := os.ReadFile("./database/sql/create_table_order.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderCustomerTable() {
	query, err := os.ReadFile("./database/sql/create_table_order_customer.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderAddressTable() {
	query, err := os.ReadFile("./database/sql/create_table_order_address.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}

func createOrderItemTable() {
	query, err := os.ReadFile("./database/sql/create_table_order_item.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}
