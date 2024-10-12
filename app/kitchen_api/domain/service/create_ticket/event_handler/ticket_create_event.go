package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type TicketCreateWhenTicketCreateHandler struct {
	orderID model.OrderID
	items   []*model.TicketItemRequest
	svc     service.CreateTicket
}

func NewTicketCreateWhenTicketCreateHandler(orderID model.OrderID, items []*model.TicketItemRequest, svc service.CreateTicket) domain_event.EventHandler {
	return &TicketCreateWhenTicketCreateHandler{svc: svc, orderID: orderID, items: items}
}

func (h *TicketCreateWhenTicketCreateHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.CommandName_TicketCreate {
		return errors.New("invalid event")
	}

	if err := h.svc.CreateTicket(ctx, h.orderID, h.items); err != nil {
		return err
	}

	return nil
}
