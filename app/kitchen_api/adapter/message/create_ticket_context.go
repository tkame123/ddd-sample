package message

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/service/create_ticket/event_handler"
	"github.com/tkame123/ddd-sample/proto/message"
)

type CreateTicketContext struct {
	strategy domain_event.EventHandler
	mes      domain_event.Message
	svc      service.CreateTicket
}

func NewCreateTicketContext(mes domain_event.Message, svc service.CreateTicket) (*CreateTicketContext, error) {
	if !event_handler.IsCreateTicketEvent(mes.Raw().Subject.Type) {
		return nil, errors.New("unknown event by CreateTicketService")
	}

	var strategy domain_event.EventHandler
	switch mes.Raw().Subject.Type {
	case message.Type_TYPE_COMMAND_TICKET_CREATE:
		var v message.CommandTicketCreate
		err := mes.Raw().Envelope.UnmarshalTo(&v)
		if err != nil {
			return nil, errors.New("invalid event")
		}
		orderId, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return nil, errors.New("invalid event")
		}
		items := make([]*model.TicketItemRequest, 0, len(v.Items))
		for _, i := range v.Items {
			itemId, err := model.ItemIdParse(i.ItemId)
			if err != nil {
				return nil, errors.New("invalid event")
			}
			items = append(items, &model.TicketItemRequest{
				ItemID:   *itemId,
				Quantity: int(i.Quantity),
			})
		}
		strategy = event_handler.NewTicketCreateWhenTicketCreateHandler(*orderId, items, svc)
	case message.Type_TYPE_COMMAND_TICKET_APPROVE:
		strategy = event_handler.NewTicketApproveWhenTicketApproveHandler(mes.ID(), svc)
	case message.Type_TYPE_COMMAND_TICKET_REJECT:
		strategy = event_handler.NewTicketRejectWhenTicketRejectHandler(mes.ID(), svc)

	default:
		return nil, errors.New("not implemented event by CreateTicketService")
	}

	return &CreateTicketContext{strategy: strategy, mes: mes, svc: svc}, nil
}

func (c *CreateTicketContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.mes.Raw())
}
