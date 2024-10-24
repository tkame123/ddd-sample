package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
)

type CreateOrderSagaState interface {
	CreateOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*create_order_saga.CreateOrderSagaState, error)
	CreateOrderSagaStateSave(ctx context.Context, state *create_order_saga.CreateOrderSagaState) error
}
