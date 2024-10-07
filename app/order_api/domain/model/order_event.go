package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderEvent interface {
	event.Event
	ID() OrderID
}

type OrderCreatedEvent struct {
	orderID OrderID
}

func NewOrderCreatedEvent(orderID OrderID) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		orderID: orderID,
	}
}

func (e *OrderCreatedEvent) Name() string {
	return "event.order.order_created"
}

func (e *OrderCreatedEvent) ID() OrderID {
	return e.orderID
}

type OrderApprovedEvent struct {
	orderID OrderID
}

func NewOrderApprovedEvent(orderID OrderID) *OrderApprovedEvent {
	return &OrderApprovedEvent{
		orderID: orderID,
	}
}

func (e *OrderApprovedEvent) Name() string {
	return "event.order.order_approved"
}

func (e *OrderApprovedEvent) ID() OrderID {
	return e.orderID
}

type OrderRejectedEvent struct {
	orderID OrderID
}

func NewOrderRejectedEvent(orderID OrderID) *OrderRejectedEvent {
	return &OrderRejectedEvent{
		orderID: orderID,
	}
}

func (e *OrderRejectedEvent) Name() string {
	return "event.order.order_rejected"
}

func (e *OrderRejectedEvent) ID() OrderID {
	return e.orderID
}
