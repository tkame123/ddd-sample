package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type Order interface {
	FindOne(ctx context.Context, id *model.OrderID) (*model.Order, error)
	Save(ctx context.Context, order *model.Order) error
}