package event

import (
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/lib/event"
)

type TicketEvent interface {
	event.Event
	ID() model.TicketID
}

type TicketCreated struct {
	ticketID model.TicketID
}

func NewTicketCreated(ticketID model.TicketID) *TicketCreated {
	return &TicketCreated{
		ticketID: ticketID,
	}
}

func (e *TicketCreated) Name() string {
	return "event.delivery.ticket_created"
}

func (e *TicketCreated) ID() model.TicketID {
	return e.ticketID
}
