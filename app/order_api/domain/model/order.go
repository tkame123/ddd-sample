package model

import (
	"errors"
)

// 集約ルート
type Order struct {
	orderID       OrderID
	approvalLimit int
	orderItems    []*OrderItem
	status        OrderStatus
}

type OrderItem struct {
	OrderID  OrderID
	SortNo   int
	ItemID   ItemID
	Price    int
	Quantity int
}

type OrderStatus = string

const (
	OrderStatus_ApprovalPending OrderStatus = "ApprovalPending"
	OrderStatus_OrderApproved   OrderStatus = "OrderApproved"
	OrderStatus_OrderRejected   OrderStatus = "OrderRejected"
)

type OrderItemRequest struct {
	Item
	quantity int
}

func NewOrder(items []*OrderItemRequest) (*Order, []OrderEvent, error) {
	// NOTE: IDの発行はInfra層で行う

	order := &Order{
		status: OrderStatus_ApprovalPending,
	}

	orderItems := make([]*OrderItem, 0, len(items))
	for i, item := range items {
		orderItems = append(orderItems, &OrderItem{
			OrderID:  order.orderID,
			SortNo:   i + 1,
			ItemID:   item.ItemID,
			Price:    item.Price,
			Quantity: item.quantity,
		})
	}
	order.orderItems = orderItems

	createdEvent := NewOrderCreatedEvent(order.OrderID())
	return order, []OrderEvent{createdEvent}, nil
}

func (o *Order) OrderID() OrderID {
	return o.orderID
}

func (o *Order) ApproveOrder() ([]OrderEvent, error) {
	if o.status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.status = OrderStatus_OrderApproved
	return []OrderEvent{NewOrderApprovedEvent(o.OrderID())}, nil
}

func (o *Order) RejectOrder() ([]OrderEvent, error) {
	if o.status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.status = OrderStatus_OrderRejected
	return []OrderEvent{NewOrderRejectedEvent(o.OrderID())}, nil
}
