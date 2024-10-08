package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) CreateOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*model.CreateOrderSagaState, error) {
	panic("implement me")
}

func (r *repo) CreateOrderSagaStateSave(ctx context.Context, order *model.CreateOrderSagaState) error {
	panic("implement me")
}
