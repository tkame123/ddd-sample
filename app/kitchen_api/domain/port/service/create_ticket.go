package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

type CreateTicket interface {
	CreateTicket(ctx context.Context, orderID model.OrderID, items []*model.TicketItemRequest) error
	ApproveTicket(ctx context.Context, orderID model.OrderID) error
	RejectTicket(ctx context.Context, orderID model.OrderID) error
}
