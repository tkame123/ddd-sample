package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type TicketRejectWhenTicketRejectHandler struct {
	orderID model.OrderID
	svc     service.CreateTicket
}

func NewTicketRejectWhenTicketRejectHandler(orderID model.OrderID, svc service.CreateTicket) domain_event.EventHandler {
	return &TicketRejectWhenTicketRejectHandler{svc: svc, orderID: orderID}
}

func (h *TicketRejectWhenTicketRejectHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_REJECT {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.svc.RejectTicket(ctx, h.orderID); err != nil {
		return err
	}

	return nil
}
