package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type order struct {
}

func NewOrder() repository.Order {
	return &order{}
}

func (o *order) FindOne(ctx context.Context, id *model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (o *order) Save(ctx context.Context, order *model.Order) error {
	// TODO: oredr -> Update Or Insert / OrderItem -> Replace
	panic("implement me")
}
