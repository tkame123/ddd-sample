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
	err := r.db.Order.Create().
		SetID(order.OrderID()).
		SetApprovalLimit(int64(order.ApprovalLimit())).
		SetStatus(fromModelStatus(order.Status())).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	// TODO: OrderItem対応

	return err
}

func fromModelStatus(status model.OrderStatus) e_order.Status {
	switch status {
	case model.OrderStatus_ApprovalPending:
		return e_order.StatusApprovalPending
	case model.OrderStatus_OrderApproved:
		return e_order.StatusOrderApproved
	case model.OrderStatus_OrderRejected:
		return e_order.StatusOrderRejected
	default:
		panic("invalid status")
	}
}
