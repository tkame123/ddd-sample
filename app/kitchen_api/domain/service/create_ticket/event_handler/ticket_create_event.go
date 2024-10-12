package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type TicketCreateWhenTicketCreateHandler struct {
	orderID model.OrderID
	items   []*model.TicketItemRequest
	svc     service.CreateTicket
}

func NewTicketCreateWhenTicketCreateHandler(orderID model.OrderID, items []*model.TicketItemRequest, svc service.CreateTicket) domain_event.EventHandler {
	return &TicketCreateWhenTicketCreateHandler{svc: svc, orderID: orderID, items: items}
}

func (h *TicketCreateWhenTicketCreateHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_CREATE {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.svc.CreateTicket(ctx, h.orderID, h.items); err != nil {
		return err
	}

	return nil
}
