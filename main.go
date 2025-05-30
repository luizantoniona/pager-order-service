package main

import (
	"log"
	"net/http"

	"pager-order-service/database"
	"pager-order-service/route"
)

func main() {
	database.InitializeDatabase()

	route.SetupRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running!"))
	})

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
