package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/proto/message"
)

type TicketRejectWhenTicketRejectHandler struct {
	svc service.CreateTicket
}

func NewTicketRejectWhenTicketRejectHandler(svc service.CreateTicket) domain_event.EventHandler {
	return &TicketRejectWhenTicketRejectHandler{svc: svc}
}

func (h *TicketRejectWhenTicketRejectHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_REJECT {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.CommandTicketReject
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	if err := h.svc.RejectTicket(ctx, *id); err != nil {
		return err
	}

	return nil
}
