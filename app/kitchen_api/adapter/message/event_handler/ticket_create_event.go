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
	svc service.Ticket
}

func NewTicketCreateWhenTicketCreateHandler(svc service.Ticket) domain_event.EventHandler {
	return &TicketCreateWhenTicketCreateHandler{svc: svc}
}

func (h *TicketCreateWhenTicketCreateHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_TICKET_CREATE {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.CommandTicketCreate
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	items := make([]*model.TicketItemRequest, 0, len(v.Items))
	for _, i := range v.Items {
		itemId, err := model.ItemIdParse(i.ItemId)
		if err != nil {
			return fmt.Errorf("failed to parse item id: %w", err)
		}
		items = append(items, &model.TicketItemRequest{
			ItemID:   *itemId,
			Quantity: int(i.Quantity),
		})
	}

	if err := h.svc.CreateTicket(ctx, *id, items); err != nil {
		return err
	}

	return nil
}
