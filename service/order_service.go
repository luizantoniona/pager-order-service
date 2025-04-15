package service

import (
	"pager-order-service/model"
	"pager-order-service/repository"
)

func InsertOrder(order model.Order) error {
	// TODO: Criar verificações

	return repository.InsertOrder(order)
}
