package model

import (
	"errors"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
)

type TicketStatus = string

const (
	Tickettatus_ApprovalPending TicketStatus = "Pending"
	Tickettatus_Approved        TicketStatus = "Approved"
	TicketStatus_Rejected       TicketStatus = "Rejected"
)

// 集約ルート
type Ticket struct {
	TicketID    TicketID      `json:"ticket_id"`
	OrderID     OrderID       `json:"order_id"`
	TicketItems []*TicketItem `json:"items"`
	Status      TicketStatus  `json:"status"`
}

type TicketItem struct {
	TicketID TicketID `json:"ticket_id"`
	ItemID   ItemID   `json:"item_id"`
	Quantity int      `json:"quantity"`
}

type TicketItemRequest struct {
	ItemID   ItemID `json:"item_id"`
	Quantity int    `json:"quantity"`
}

// TODO: 重複オーダチェックの際に、TicketCreationFailedEventを発行する
func NewTicket(orderID OrderID, items []*TicketItemRequest) (*Ticket, []*message.Message, error) {
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

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_TICKET_CREATED,
		message.Service_SERVICE_KITCHEN,
		&message.EventTicketCreated{
			OrderId:  ticket.OrderID.String(),
			TicketId: ticket.TicketID.String(),
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return ticket, []*message.Message{mes}, nil
}

func (t *Ticket) ApproveTicket() ([]*message.Message, error) {
	if t.Status != Tickettatus_ApprovalPending {
		return nil, errors.New("ticket is not approval pending status")
	}

	t.Status = Tickettatus_Approved

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_ORDER_APPROVED,
		message.Service_SERVICE_KITCHEN,
		&message.EventTicketApproved{
			OrderId:  t.OrderID.String(),
			TicketId: t.TicketID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (t *Ticket) RejectTicket() ([]*message.Message, error) {
	if t.Status != Tickettatus_ApprovalPending {
		return nil, errors.New("ticket is not approval pending status")
	}

	t.Status = TicketStatus_Rejected

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_ORDER_REJECTED,
		message.Service_SERVICE_KITCHEN,
		&message.EventTicketRejected{
			OrderId:  t.OrderID.String(),
			TicketId: t.TicketID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}
