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
	svc service.Ticket
}

func NewTicketApproveWhenTicketApproveHandler(svc service.Ticket) domain_event.EventHandler {
	return &TicketApproveWhenTicketApproveHandler{svc: svc}
}

func (h *TicketApproveWhenTicketApproveHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_APPROVE {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.CommandTicketApprove
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	ticketId, err := model.TicketIdParse(v.TicketId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	if err := h.svc.ApproveTicket(ctx, *id, *ticketId); err != nil {
		return err
	}

	return nil
}
