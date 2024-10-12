package message

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/service/create_ticket/event_handler"
	"github.com/tkame123/ddd-sample/proto/message"
)

type CreateTicketContext struct {
	strategy domain_event.EventHandler
	mes      *message.Message
	svc      service.CreateTicket
}

func NewCreateTicketContext(mes *message.Message, svc service.CreateTicket) (*CreateTicketContext, error) {
	if !event_handler.IsCreateTicketEvent(mes.Subject.Type) {
		return nil, errors.New("unknown event by CreateTicketService")
	}

	return &CreateTicketContext{
		strategy: event_handler.EventMap[mes.Subject.Type](svc),
		mes:      mes,
		svc:      svc,
	}, nil
}

func (c *CreateTicketContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.mes)
}
