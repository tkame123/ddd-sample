package model

import (
	"errors"
	"github.com/tkame123/ddd-sample/lib/event"
)

const APPOVAL_LIMIT = 10000

// 集約ルート
type Order struct {
	OrderID    OrderID
	OrderItems []*OrderItem
	Status     OrderStatus
}

type OrderItem struct {
	OrderItemID OrderItemID
	OrderID     OrderID
	SortNo      int
	ItemID      ItemID
	Price       int
	Quantity    int
}

type OrderStatus = string

const (
	OrderStatus_ApprovalPending OrderStatus = "ApprovalPending"
	OrderStatus_OrderApproved   OrderStatus = "OrderApproved"
	OrderStatus_OrderRejected   OrderStatus = "OrderRejected"
)

type OrderItemRequest struct {
	Item
	Quantity int
}

func NewOrder(items []*OrderItemRequest) (*Order, []event.Event, error) {
	orderID := generateID()

	orderItems := make([]*OrderItem, 0, len(items))
	for i, item := range items {
		orderItems = append(orderItems, &OrderItem{
			OrderItemID: generateID(),
			OrderID:     orderID,
			SortNo:      i + 1,
			ItemID:      item.ItemID,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	order := &Order{
		OrderID:    orderID,
		OrderItems: orderItems,
		Status:     OrderStatus_ApprovalPending,
	}

	if !order.validateApprovalLimit() {
		return nil, nil, errors.New("approval limit over")
	}

	return order, []event.Event{&OrderCreatedEvent{OrderID: order.OrderID}}, nil
}

func (o *Order) ApproveOrder() ([]event.Event, error) {
	if o.Status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.Status = OrderStatus_OrderApproved
	return []event.Event{&OrderApprovedEvent{o.OrderID}}, nil
}

func (o *Order) RejectOrder() ([]event.Event, error) {
	if o.Status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.Status = OrderStatus_OrderRejected
	return []event.Event{&OrderRejectedEvent{o.OrderID}}, nil
}

func (o *Order) validateApprovalLimit() bool {
	sum := 0
	for _, v := range o.OrderItems {
		sum += v.Price * v.Quantity
	}

	return sum <= APPOVAL_LIMIT
}
