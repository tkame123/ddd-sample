package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type TicketApproveWhenTicketApproveHandler struct {
	orderID model.OrderID
	svc     service.CreateTicket
}

func NewTicketApproveWhenTicketApproveHandler(orderID model.OrderID, svc service.CreateTicket) domain_event.EventHandler {
	return &TicketApproveWhenTicketApproveHandler{svc: svc, orderID: orderID}
}

func (h *TicketApproveWhenTicketApproveHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_APPROVE {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.svc.ApproveTicket(ctx, h.orderID); err != nil {
		return err
	}

	return nil
}
