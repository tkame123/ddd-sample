package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	e_order "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	e_orderItem "github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/orderitem"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"time"
)

func (r *repo) OrderFindOne(ctx context.Context, id model.OrderID) (*model.Order, error) {
	order, err := r.db.Order.Query().
		Where(e_order.ID(id)).
		WithOrderItems().
		First(ctx)
	if err != nil {
		return nil, err
	}

	return toModelOrder(order), nil
}

func (r *repo) OrderSave(ctx context.Context, order *model.Order) error {
	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		prev := int64(order.Version)
		next := time.Now().UnixNano()

		exits, err := tx.Order.Query().
			Where(e_order.ID(order.OrderID), e_order.Version(prev)).
			Exist(ctx)
		if err != nil {
			return err
		}

		if exits {
			_, err = tx.Order.Update().
				Where(e_order.ID(order.OrderID), e_order.Version(prev)).
				SetStatus(fromModelOrderStatus(order.Status)).
				SetVersion(next).
				Save(ctx)
		} else {
			_, err = tx.Order.Create().
				SetID(order.OrderID).
				SetStatus(fromModelOrderStatus(order.Status)).
				Save(ctx)
		}
		if err != nil {
			return err
		}

		_, err = tx.OrderItem.Delete().
			Where(e_orderItem.OrderID(order.OrderID)).
			Exec(ctx)
		if err != nil {
			return err
		}

		orderItems := order.OrderItems
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

func toModelOrder(order *ent.Order) *model.Order {
	items := make([]*model.OrderItem, 0, len(order.Edges.OrderItems))

	for _, item := range order.Edges.OrderItems {
		items = append(items, toModelOrderItem(item))
	}

	return &model.Order{
		OrderID:    order.ID,
		Status:     toModelOrderStatus(order.Status),
		OrderItems: items,
		Version:    int(order.Version),
	}
}

func toModelOrderItem(orderItem *ent.OrderItem) *model.OrderItem {
	return &model.OrderItem{
		OrderItemID: orderItem.ID,
		OrderID:     orderItem.OrderID,
		SortNo:      int(orderItem.SortNo),
		ItemID:      orderItem.ItemID,
		Price:       int(orderItem.Price),
		Quantity:    int(orderItem.Quantity),
	}
}

func fromModelOrderStatus(status model.OrderStatus) e_order.Status {
	switch status {
	case model.OrderStatus_ApprovalPending:
		return e_order.StatusPending
	case model.OrderStatus_Approved:
		return e_order.StatusApproved
	case model.OrderStatus_Rejected:
		return e_order.StatusRejected
	case model.OrderStatus_CancelPending:
		return e_order.StatusCancelPending
	case model.OrderStatus_Canceled:
		return e_order.StatusCanceled
	default:
		panic("invalid status")
	}
}

func toModelOrderStatus(status e_order.Status) model.OrderStatus {
	switch status {
	case e_order.StatusPending:
		return model.OrderStatus_ApprovalPending
	case e_order.StatusApproved:
		return model.OrderStatus_Approved
	case e_order.StatusRejected:
		return model.OrderStatus_Rejected
	case e_order.StatusCancelPending:
		return model.OrderStatus_CancelPending
	case e_order.StatusCanceled:
		return model.OrderStatus_Canceled
	default:
		panic("invalid status")
	}
}
