package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

type Ticket interface {
	TicketFindOne(ctx context.Context, id *model.TicketID) (*model.Ticket, error)
	TicketSave(ctx context.Context, order *model.Ticket) error
}
