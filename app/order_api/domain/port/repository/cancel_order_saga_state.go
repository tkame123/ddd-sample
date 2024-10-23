package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
)

type CancelOrderSagaState interface {
	CancelOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*cancel_order_saga.CancelOrderSagaState, error)
	CancelOrderSagaStateSave(ctx context.Context, state *cancel_order_saga.CancelOrderSagaState) error
}
