package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type Order interface {
	OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error)
	OrderSave(ctx context.Context, order *model.Order) error
}
