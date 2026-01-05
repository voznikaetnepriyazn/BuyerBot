package csharp

import (
	"context"
	"telegram/internal/model/customer"
)

type CustomerGoodGateway interface {
	AddCustomer(ctx context.Context, Id string) (customer.Customer, error)
	Deletecustomer(ctx context.Context, Id string) error
	GetAllCustomers(ctx context.Context) ([]customer.Customer, error)
	GetByIdCustomer(ctx context.Context, Id string) (customer.Customer, error)
	UpdateCustomer(ctx context.Context, Id string) error
	IsCustomerCreated(ctx context.Context, Id string) (bool, error)
}
