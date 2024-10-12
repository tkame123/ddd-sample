package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

type CreateTicket interface {
	CreateTicket(ctx context.Context, items []*model.TicketItemRequest) (model.TicketID, error)
	ApproveTicket(ctx context.Context, ticketID model.TicketID) (model.TicketID, error)
	RejectTicket(ctx context.Context, ticketID model.TicketID) (model.TicketID, error)
}
