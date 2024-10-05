package event

import (
	"github.com/tkame123/ddd-sample/app/lib/event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type OrderEvent interface {
	event.Event
	ID() model.OrderID
}

type OrderCreated struct {
	orderID model.OrderID
}

func NewOrderCreated(orderID model.OrderID) *OrderCreated {
	return &OrderCreated{
		orderID: orderID,
	}
}

func (e *OrderCreated) Name() string {
	return "event.order.created"
}

func (e *OrderCreated) ID() model.OrderID {
	return e.orderID
}

type OrderItemsUpdated struct {
	orderID model.OrderID
}

func NewOrderItemsUpdated(orderID model.OrderID) *OrderItemsUpdated {
	return &OrderItemsUpdated{
		orderID: orderID,
	}
}

func (e *OrderItemsUpdated) Name() string {
	return "event.order.items_updated"
}

func (e *OrderItemsUpdated) ID() model.OrderID {
	return e.orderID
}
