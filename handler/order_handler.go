package handler

import (
	"encoding/json"
	"net/http"

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
	orderID := r.URL.Query().Get("id")

	if orderID == "" {
		orderIDs, err := service.GetAllOrderIDs()
		if err != nil {
			http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orderIDs)

		return
	}

	pager, err := service.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pager)
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
