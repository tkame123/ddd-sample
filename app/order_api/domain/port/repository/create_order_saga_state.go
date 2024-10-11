package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrderSagaState interface {
	CreateOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*model.CreateOrderSagaState, error)
	CreateOrderSagaStateSave(ctx context.Context, state *model.CreateOrderSagaState) error
}
