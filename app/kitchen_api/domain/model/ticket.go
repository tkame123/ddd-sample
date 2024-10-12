package model

import (
	"errors"
	"github.com/tkame123/ddd-sample/lib/event"
)

type TicketStatus = string

const (
	Tickettatus_ApprovalPending TicketStatus = "Pending"
	Tickettatus_Approved        TicketStatus = "Approved"
	TicketStatus_Rejected       TicketStatus = "Rejected"
)

// 集約ルート
type Ticket struct {
	TicketID    TicketID
	OrderID     OrderID
	TicketItems []*TicketItem
	Status      TicketStatus
}

type TicketItem struct {
	TicketID TicketID
	ItemID   ItemID
	Quantity int
}

type TicketItemRequest struct {
	ItemID   ItemID
	Quantity int
}

// TODO: 重複オーダチェックの際に、TicketCreationFailedEventを発行する
func NewTicket(orderID OrderID, items []*TicketItemRequest) (*Ticket, []event.Event, error) {
	ticketID := generateID()

	ticketItems := make([]*TicketItem, 0, len(items))
	for _, item := range items {
		ticketItems = append(ticketItems, &TicketItem{
			TicketID: ticketID,
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		})
	}

	ticket := &Ticket{
		TicketID:    ticketID,
		OrderID:     orderID,
		TicketItems: ticketItems,
	}

	return ticket, []event.Event{&TicketCreatedEvent{TicketID: ticket.TicketID}}, nil
}

func (t *Ticket) ApproveTicket() ([]event.Event, error) {
	if t.Status != Tickettatus_ApprovalPending {
		return nil, errors.New("ticket is not approval pending status")
	}

	return []event.Event{&TicketApprovedEvent{TicketID: t.TicketID}}, nil
}

func (t *Ticket) RejectTicket() ([]event.Event, error) {
	if t.Status != Tickettatus_ApprovalPending {
		return nil, errors.New("ticket is not approval pending status")
	}

	return []event.Event{&TicketRejectedEvent{TicketID: t.TicketID}}, nil
}
