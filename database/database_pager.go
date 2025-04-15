package database

import (
	"log"
	"os"
)

func createPagerTable() {
	query, err := os.ReadFile("./database/sql/schema/create_table_pager.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}

func createPagerCustomerTable() {
	query, err := os.ReadFile("./database/sql/schema/create_table_pager_customer.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}

func createPagerItemTable() {
	query, err := os.ReadFile("./database/sql/schema/create_table_pager_item.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}
