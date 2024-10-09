package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	e_order "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	e_orderItem "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/orderitem"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (r *repo) OrderSave(ctx context.Context, order *model.Order) error {
	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		err := tx.Order.Create().
			SetID(order.OrderID()).
			SetApprovalLimit(int64(order.ApprovalLimit())).
			SetStatus(fromModelStatus(order.Status())).
			OnConflictColumns("id").
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.OrderItem.Delete().
			Where(e_orderItem.OrderID(order.OrderID())).
			Exec(ctx)
		if err != nil {
			return err
		}

		orderItems := order.OrderItems()
		err = tx.OrderItem.MapCreateBulk(
			orderItems,
			func(
				oi *ent.OrderItemCreate, i int) {
				item := orderItems[i]
				oi.SetID(item.OrderItemID)
				oi.SetOrderID(item.OrderID)
				oi.SetItemID(item.ItemID)
				oi.SetSortNo(int32(item.SortNo))
				oi.SetPrice(int64(item.Price))
				oi.SetQuantity(int32(item.Quantity))
			},
		).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
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
