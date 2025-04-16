package service

import (
	"pager-order-service/model"
	"pager-order-service/repository"
)

func GetOrderByID(id string) (model.Order, error) {
	return repository.GetOrderByID(id)
}

func GetAllOrderIDs() ([]string, error) {
	return repository.GetAllOrderIDs()
}

func InsertOrder(order model.Order) error {
	// TODO: Criar verificações

	return repository.InsertOrder(order)
}
