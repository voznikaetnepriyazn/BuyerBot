package services

import (
	"Order/internal/storage"
)

type OrderStruct struct {
	storage storage.StorageInter
}
type OrderService interface {
	IsOrderCreated(id int) bool
}

func (o *OrderStruct) IsOrderCreated(id int64) (bool, error) {
	order, err := o.storage.GetByIdURL(id)
	if err != nil {
		return false, err
	}

	if order == 0 {
		return false, err
	}

	return true, nil
}
