package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderEvent interface {
	event.Event
	ID() OrderID
}

type OrderCreatedEvent struct {
	OrderID OrderID
}

func NewOrderCreatedEvent(orderID OrderID) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		OrderID: orderID,
	}
}

func (e *OrderCreatedEvent) Name() event.Name {
	return event.EventName_OrderCreated
}

func (e *OrderCreatedEvent) ID() OrderID {
	return e.OrderID
}

type OrderApprovedEvent struct {
	OrderID OrderID
}

func NewOrderApprovedEvent(orderID OrderID) *OrderApprovedEvent {
	return &OrderApprovedEvent{
		OrderID: orderID,
	}
}

func (e *OrderApprovedEvent) Name() event.Name {
	return event.EventName_OrderApproved
}

func (e *OrderApprovedEvent) ID() OrderID {
	return e.OrderID
}

type OrderRejectedEvent struct {
	OrderID OrderID
}

func NewOrderRejectedEvent(orderID OrderID) *OrderRejectedEvent {
	return &OrderRejectedEvent{
		OrderID: orderID,
	}
}

func (e *OrderRejectedEvent) Name() event.Name {
	return event.EventName_OrderRejected
}

func (e *OrderRejectedEvent) ID() OrderID {
	return e.OrderID
}
