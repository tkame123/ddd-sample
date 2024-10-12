package message

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/service/create_ticket/event_handler"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type CreateTicketContext struct {
	strategy domain_event.EventHandler
	event    event_helper.Event
	svc      service.CreateTicket
}

func NewCreateTicketContext(ev event_helper.Event, svc service.CreateTicket) (*CreateTicketContext, error) {
	if !event_handler.IsCreateTicketEvent(ev.Name()) {
		return nil, errors.New("unknown event by CreateTicketService")
	}

	var strategy domain_event.EventHandler
	switch ev.Name() {
	case event_helper.CommandName_TicketCreate:
		v, ok := ev.(*model.TicketCreateCommand)
		if !ok {
			return nil, errors.New("invalid event")
		}
		strategy = event_handler.NewTicketCreateWhenTicketCreateHandler(v.OrderID, v.Items, svc)
	case event_helper.CommandName_TicketApprove:
		strategy = event_handler.NewTicketApproveWhenTicketApproveHandler(ev.ID(), svc)
	case event_helper.CommandName_TicketReject:
		strategy = event_handler.NewTicketRejectWhenTicketRejectHandler(ev.ID(), svc)

	default:
		return nil, errors.New("not implemented event by CreateTicketService")
	}

	return &CreateTicketContext{strategy: strategy, event: ev, svc: svc}, nil
}

func (c *CreateTicketContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.event)
}
