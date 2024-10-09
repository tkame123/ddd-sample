package database

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	e_order "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	e_orderItem "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/orderitem"

	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error) {
	panic("implement me")
}

func (r *repo) OrderSave(ctx context.Context, order *model.Order) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}
	err = tx.Order.Create().
		SetID(order.OrderID()).
		SetApprovalLimit(int64(order.ApprovalLimit())).
		SetStatus(fromModelStatus(order.Status())).
		OnConflictColumns("id").
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.OrderItem.Delete().
		Where(e_orderItem.OrderID(order.OrderID())).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
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
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, err)
	}

	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
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
