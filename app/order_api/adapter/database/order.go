package database

import (
	"context"
	e_order "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (r *repo) OrderSave(ctx context.Context, order *model.Order) error {
	// TODO: oredr -> Update Or Insert / OrderItem -> Replace

	_, err := r.db.Order.Create().
		SetApprovalLimit(1000).
		SetStatus(e_order.StatusApprovalPending).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
