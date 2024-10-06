package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/repository"
)

type CreateTicket struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewCreateTicket(rep repository.Repository, pub domain_event.Publisher) *CreateTicket {
	return &CreateTicket{
		rep: rep,
		pub: pub,
	}
}

type CreateTicketInput struct {
	OrderID model.OrderID
	Items   []*model.TicketItemRequest
}

type CreateTicketOutput struct {
	TicketID model.TicketID
}

func (c *CreateTicket) Execute(ctx context.Context, input CreateTicketInput) (*CreateTicketOutput, error) {
	ticket, events, err := model.NewTicket(input.OrderID, input.Items)
	if err != nil {
		return nil, err
	}

	if err := c.rep.Ticket.Save(ctx, ticket); err != nil {
		return nil, err
	}

	c.pub.PublishMessages(ctx, events)

	return &CreateTicketOutput{TicketID: ticket.TicketID()}, nil
}
