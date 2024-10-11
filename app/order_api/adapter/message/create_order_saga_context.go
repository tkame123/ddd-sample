package message

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/lib/event"
)

type CreateOrderSagaContext struct {
	rep      repository.Repository
	strategy domain_event.EventHandler
	event    event.Event
}

func NewCreateOrderSagaContext(ev event.Event) (*CreateOrderSagaContext, error) {
	if !event_handler.IsCreateOrderSagaEvent(ev.Name()) {
		return nil, errors.New("unknown event by CreateOrderSagaContext")
	}

	// TODO: Sagaの復元

	var strategy domain_event.EventHandler
	switch ev.Name() {
	case event.EventName_OrderCreated:

	default:
		return nil, errors.New("not implemented event by CreateOrderSagaContext")
	}

	return &CreateOrderSagaContext{strategy: strategy}, nil
}

func (c *CreateOrderSagaContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.event)
}
