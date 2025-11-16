/*public async Task<bool> IsOrderCreated(Guid Id)
{
    Order order = await orderRepository.GetByIdAsync(Id);
    if (order == null)
    {
        return false;
    }
    return true;
}*/

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
