package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type TicketRejectWhenTicketRejectHandler struct {
	orderID model.OrderID
	svc     service.CreateTicket
}

func NewTicketRejectWhenTicketRejectHandler(orderID model.OrderID, svc service.CreateTicket) domain_event.EventHandler {
	return &TicketRejectWhenTicketRejectHandler{svc: svc, orderID: orderID}
}

func (h *TicketRejectWhenTicketRejectHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.CommandName_TicketReject {
		return errors.New("invalid event")
	}

	if err := h.svc.RejectTicket(ctx, h.orderID); err != nil {
		return err
	}

	return nil
}
