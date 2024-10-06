package model

import "github.com/tkame123/ddd-sample/app/kitchen_api/domain/event"

// 集約ルート
type Ticket struct {
	ticketID    TicketID
	orderID     OrderID
	ticketItems []*TicketItem
}

type TicketItem struct {
	ticketID TicketID
	ItemID   ItemID
	Quantity int64
}

type TicketItemRequest struct {
	ItemID   ItemID
	quantity int64
}

func NewTicket(orderID OrderID, items []*TicketItemRequest) (*Ticket, []event.TicketEvent, error) {
	ticketID := generateID()

	ticketItems := make([]*TicketItem, 0, len(items))
	for _, item := range items {
		ticketItems = append(ticketItems, &TicketItem{
			ticketID: ticketID,
			ItemID:   item.ItemID,
			Quantity: item.quantity,
		})
	}

	ticket := &Ticket{
		ticketID:    ticketID,
		orderID:     orderID,
		ticketItems: ticketItems,
	}

	createEvent := event.NewTicketCreated(ticket.ticketID)

	return ticket, []event.TicketEvent{createEvent}, nil
}

func (s *Ticket) TicketID() TicketID {
	return s.ticketID
}
