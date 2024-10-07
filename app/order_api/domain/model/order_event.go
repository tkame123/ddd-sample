package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderEvent interface {
	event.Event
	ID() OrderID
}

type ApproveOrderCommand struct {
	orderID OrderID
}

func NewApproveOrderCommand(orderID OrderID) *ApproveOrderCommand {
	return &ApproveOrderCommand{
		orderID: orderID,
	}
}

func (e *ApproveOrderCommand) Name() string {
	return "command.order.approve_order"
}

func (e *ApproveOrderCommand) ID() OrderID {
	return e.orderID
}

type RejectOrderCommand struct {
	orderID OrderID
}

func NewRejectOrderCommand(orderID OrderID) *RejectOrderCommand {
	return &RejectOrderCommand{
		orderID: orderID,
	}
}

func (e *RejectOrderCommand) Name() string {
	return "command.order.reject_order"
}

func (e *RejectOrderCommand) ID() OrderID {
	return e.orderID
}

// message

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
