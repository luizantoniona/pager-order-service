package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {

	createOrderTable()
	createOrderCustomerTable()
	createOrderAddressTable()
	createOrderItemTable()

	createPagerTable()
	createPagerCustomerTable()
	createPagerItemTable()

	fmt.Println("Database and tables created successfully!")
}
