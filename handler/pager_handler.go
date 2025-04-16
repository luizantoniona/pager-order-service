package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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
	if r.URL.Path == "/pagers" {
		pagerIDs, err := service.GetAllPagerIDs()
		if err != nil {
			http.Error(w, "Failed to retrieve pagers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pagerIDs)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/pagers/") {
		id := strings.TrimPrefix(r.URL.Path, "/pagers/")
		pager, err := service.GetPagerByID(id)
		if err != nil {
			http.Error(w, "Pager not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pager)
		return
	}

	http.NotFound(w, r)
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
