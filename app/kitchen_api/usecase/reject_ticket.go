package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/repository"
)

type RejectTicket struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewRejectTicket(rep repository.Repository, pub domain_event.Publisher) *RejectTicket {
	return &RejectTicket{
		rep: rep,
		pub: pub,
	}
}

type RejectTicketInput struct {
	OrderID model.OrderID
}

type RejectTicketOutput struct {
	TicketID model.TicketID
}

func (c *RejectTicket) Execute(ctx context.Context, input RejectTicketInput) (*RejectTicketOutput, error) {
	panic("implement me")
}
