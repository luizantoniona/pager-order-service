package handler

import (
	"encoding/json"
	"net/http"

	"pager-order-service/model"
	"pager-order-service/service"
)

func HandlePagers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreatePagerHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
