package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrderSagaState interface {
	FindOne(ctx context.Context, id model.OrderID) (*model.CreateOrderSagaState, error)
	Save(ctx context.Context, order *model.CreateOrderSagaState) error
}
