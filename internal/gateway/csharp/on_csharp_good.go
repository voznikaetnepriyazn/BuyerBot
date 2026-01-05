package csharp

import (
	"context"
	"telegram/internal/model/good"
)

type GoodGateway interface {
	AddGood(ctx context.Context, Id string) (good.Good, error)
	DeleteGood(ctx context.Context, Id string) error
	GetAllGoods(ctx context.Context) ([]good.Good, error)
	GetByIdGood(ctx context.Context, Id string) (good.Good, error)
	UpdateGood(ctx context.Context, Id string) error
	IsAvaliableForOrder(ctx context.Context, Id string) (bool, error)
	RestOfGood(ctx context.Context, Id string) (int64, error)
}
