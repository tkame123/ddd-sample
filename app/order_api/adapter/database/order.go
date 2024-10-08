package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (r *repo) OrderSave(ctx context.Context, order *model.Order) error {
	// TODO: oredr -> Update Or Insert / OrderItem -> Replace
	panic("implement me")
}
