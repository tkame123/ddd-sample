package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderCreatedEvent struct {
	OrderID OrderID
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

func (e *OrderApprovedEvent) Name() event.Name {
	return event.EventName_OrderApproved
}

func (e *OrderApprovedEvent) ID() OrderID {
	return e.OrderID
}

type OrderRejectedEvent struct {
	OrderID OrderID
}

func (e *OrderRejectedEvent) Name() event.Name {
	return event.EventName_OrderRejected
}

func (e *OrderRejectedEvent) ID() OrderID {
	return e.OrderID
}
