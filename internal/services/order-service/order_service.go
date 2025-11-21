package services

import (
	"Order/internal/storage"
)

type OrderStruct struct {
	storage storage.StorageInter
}
type OrderService interface {
	IsOrderCreated(id string) bool
}

func NewService(storage storage.StorageInter) *OrderStruct {
	return &OrderStruct{
		storage: storage,
	}
}

func (o *OrderStruct) IsOrderCreated(id string) (bool, error) {
	order, err := o.storage.GetByIdURL(id)
	if err != nil {
		return false, err
	}

	if order == "" {
		return false, err
	}

	return true, nil
}
