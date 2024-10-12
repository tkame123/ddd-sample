package message

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type CreateOrderSagaContext struct {
	strategy domain_event.EventHandler
	event    event_helper.Event
}

func NewCreateOrderSagaContext(ev event_helper.Event, saga *create_order_saga.CreateOrderSaga) (*CreateOrderSagaContext, error) {
	if !event_handler.IsCreateOrderSagaEvent(ev.Name()) {
		return nil, errors.New("unknown event by CreateOrderSagaContext")
	}

	var strategy domain_event.EventHandler
	switch ev.Name() {
	case event_helper.EventName_OrderCreated:
		strategy = event_handler.NewNextStepSagaWhenOrderCreatedHandler(saga)
	case event_helper.EventName_OrderApproved:
		strategy = event_handler.NewNextStepSagaWhenOrderApprovedHandler(saga)
	case event_helper.EventName_OrderRejected:
		strategy = event_handler.NewNextStepSagaWhenOrderRejectedHandler(saga)
	case event_helper.EventName_TicketCreated:
		strategy = event_handler.NewNextStepSagaWhenTicketCreatedHandler(saga)
	case event_helper.EventName_TicketCreationFailed:
		strategy = event_handler.NewNextStepSagaWhenTicketCreationFailedHandler(saga)
	case event_helper.EventName_TicketApproved:
		strategy = event_handler.NewNextStepSagaWhenTicketApprovedHandler(saga)
	case event_helper.EventName_TicketRejected:
		strategy = event_handler.NewNextStepSagaWhenTicketRejectedHandler(saga)
	case event_helper.EventName_CardAuthorized:
		strategy = event_handler.NewNextStepSagaWhenCardAuthorizedHandler(saga)
	case event_helper.EventName_CardAuthorizeFailed:
		strategy = event_handler.NewNextStepSagaWhenCardAuthorizeFailedHandler(saga)

	default:
		return nil, errors.New("not implemented event by CreateOrderSagaContext")
	}

	return &CreateOrderSagaContext{strategy: strategy, event: ev}, nil
}

func (c *CreateOrderSagaContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.event)
}
