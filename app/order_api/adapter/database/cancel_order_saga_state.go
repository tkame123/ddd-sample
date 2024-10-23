package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
)

func (r *repo) CancelOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*cancel_order_saga.CancelOrderSagaState, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) CancelOrderSagaStateSave(ctx context.Context, state *cancel_order_saga.CancelOrderSagaState) error {
	//TODO implement me
	panic("implement me")
}
