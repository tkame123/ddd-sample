package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type createOrderSagaState struct {
}

func NewCreateOrderSagaState() repository.CreateOrderSagaState {
	return &createOrderSagaState{}
}

func (o *createOrderSagaState) FindOne(ctx context.Context, id *model.OrderID) (*model.CreateOrderSagaState, error) {
	panic("implement me")
}

func (o *createOrderSagaState) Save(ctx context.Context, order *model.CreateOrderSagaState) error {
	panic("implement me")
}
