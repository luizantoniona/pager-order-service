package service

import (
	"pager-order-service/model"
	"pager-order-service/repository"
)

func InsertPager(pager model.Pager) error {
	// TODO: Criar verificações

	return repository.InsertPager(pager)
}
