package model

import (
	"github.com/tkame123/ddd-sample/lib/event"
)

type TicketEvent interface {
	event.Event
	ID() TicketID
}

type TicketCreated struct {
	ticketID TicketID
}

func NewTicketCreated(ticketID TicketID) *TicketCreated {
	return &TicketCreated{
		ticketID: ticketID,
	}
}

func (e *TicketCreated) Name() string {
	return "event.delivery.ticket_created"
}

func (e *TicketCreated) ID() TicketID {
	return e.ticketID
}

type TicketRejected struct {
	ticketID TicketID
}

func NewTTicketRejected(ticketID TicketID) *TicketRejected {
	return &TicketRejected{
		ticketID: ticketID,
	}
}

func (e *TicketRejected) Name() string {
	return "event.delivery.ticket_rejected"
}

func (e *TicketRejected) ID() TicketID {
	return e.ticketID
}
