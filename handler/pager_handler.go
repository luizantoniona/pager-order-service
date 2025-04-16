package handler

import (
	"encoding/json"
	"net/http"

	"pager-order-service/model"
	"pager-order-service/service"
)

func HandlePagers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetPagerHandler(w, r)
	case http.MethodPost:
		CreatePagerHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetPagerHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pagers/"):]
	if r.URL.Path != "/pagers" && id != "" {
		order, err := service.GetPagerByID(id)
		if err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
		return
	}

	pagerIDs, err := service.GetAllPagerIDs()
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagerIDs)
}

func CreatePagerHandler(w http.ResponseWriter, r *http.Request) {
	var pager model.Pager

	err := json.NewDecoder(r.Body).Decode(&pager)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = service.InsertPager(pager)
	if err != nil {
		http.Error(w, "Failed to create pager", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pager)
}
