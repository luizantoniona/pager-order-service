package service

import (
	"pager-order-service/model"
	"pager-order-service/repository"
)

func GetPagerByID(id string) (model.Pager, error) {
	return repository.GetPagerByID(id)
}

func GetAllPagerIDs() ([]string, error) {
	return repository.GetAllPagerIDs()
}

func InsertPager(pager model.Pager) error {
	// TODO: Criar verificações

	return repository.InsertPager(pager)
}
