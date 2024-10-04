package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) FindOne(ctx context.Context, id *model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (o *Order) Save(ctx context.Context, order *model.Order) error {
	// TODO: oredr -> Update Or Insert / OrderItem -> Replace
	panic("implement me")
}
