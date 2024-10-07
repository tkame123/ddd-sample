package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type TicketEvent interface {
	event.Event
	ID() TicketID
}

type TicketCreatedEvent struct {
	ticketID TicketID
}

func NewTicketCreatedEvent(ticketID TicketID) *TicketCreatedEvent {
	return &TicketCreatedEvent{
		ticketID: ticketID,
	}
}

func (e *TicketCreatedEvent) Name() string {
	return "event.delivery.ticket_created"
}

func (e *TicketCreatedEvent) ID() TicketID {
	return e.ticketID
}

type TicketRejectedEvent struct {
	ticketID TicketID
}

func NewTTicketRejectedEvent(ticketID TicketID) *TicketRejectedEvent {
	return &TicketRejectedEvent{
		ticketID: ticketID,
	}
}

func (e *TicketRejectedEvent) Name() string {
	return "event.delivery.ticket_rejected"
}

func (e *TicketRejectedEvent) ID() TicketID {
	return e.ticketID
}
