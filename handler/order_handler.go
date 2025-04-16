package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"pager-order-service/model"
	"pager-order-service/service"
)

func HandleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetOrderHandler(w, r)
	case http.MethodPost:
		CreateOrderHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/orders" {
		orderIDs, err := service.GetAllOrderIDs()
		if err != nil {
			http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orderIDs)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/orders/") {
		id := strings.TrimPrefix(r.URL.Path, "/orders/")
		order, err := service.GetOrderByID(id)
		if err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
		return
	}

	http.NotFound(w, r)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = service.InsertOrder(order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
