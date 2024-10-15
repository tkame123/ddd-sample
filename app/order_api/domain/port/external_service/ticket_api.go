package external_service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type KitchenAPI interface {
	CreateTicket(ctx context.Context, orderID model.OrderID) error
	ApproveTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error
	RejectTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error
}
