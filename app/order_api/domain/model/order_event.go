package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderEvent interface {
	event.Event
	ID() OrderID
}

type OrderCreated struct {
	orderID OrderID
}

func NewOrderCreated(orderID OrderID) *OrderCreated {
	return &OrderCreated{
		orderID: orderID,
	}
}

func (e *OrderCreated) Name() string {
	return "event.order.created"
}

func (e *OrderCreated) ID() OrderID {
	return e.orderID
}

type OrderItemsUpdated struct {
	orderID OrderID
}
