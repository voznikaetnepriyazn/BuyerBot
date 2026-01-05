package goo

import (
	"context"
	"telegram/internal/model/order"
)

type OrderGateway interface {
	AddOrder(ctx context.Context, Id string) (order.Order, error)
	DeleteOrder(ctx context.Context, Id string) error
	GetAllOrders(ctx context.Context) ([]order.Order, error)
	GetByIdOrder(ctx context.Context, Id string) (order.Order, error)
	UpdateOrder(ctx context.Context, Id string) error
	IsOrderCreated(ctx context.Context, Id string) (bool, error)
}
