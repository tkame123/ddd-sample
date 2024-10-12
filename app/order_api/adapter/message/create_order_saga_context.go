package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/proto/message"
)

type CreateOrderSagaContext struct {
	strategy domain_event.EventHandler
	mes      *message.Message
}

func NewCreateOrderSagaContext(mes *message.Message, saga *create_order_saga.CreateOrderSaga) (*CreateOrderSagaContext, error) {
	if !event_handler.IsCreateOrderSagaEvent(mes.Subject.Type) {
		return nil, fmt.Errorf("invalid event type: %s", mes.Subject.Type)
	}

	var strategy domain_event.EventHandler
	switch mes.Subject.Type {
	case message.Type_TYPE_EVENT_ORDER_CREATED:
		strategy = event_handler.NewNextStepSagaWhenOrderCreatedHandler(saga)
	case message.Type_TYPE_EVENT_ORDER_APPROVED:
		strategy = event_handler.NewNextStepSagaWhenOrderApprovedHandler(saga)
	case message.Type_TYPE_EVENT_ORDER_REJECTED:
		strategy = event_handler.NewNextStepSagaWhenOrderRejectedHandler(saga)
	case message.Type_TYPE_EVENT_TICKET_CREATED:
		strategy = event_handler.NewNextStepSagaWhenTicketCreatedHandler(saga)
	case message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:
		strategy = event_handler.NewNextStepSagaWhenTicketCreationFailedHandler(saga)
	case message.Type_TYPE_EVENT_TICKET_APPROVED:
		strategy = event_handler.NewNextStepSagaWhenTicketApprovedHandler(saga)
	case message.Type_TYPE_EVENT_TICKET_REJECTED:
		strategy = event_handler.NewNextStepSagaWhenTicketRejectedHandler(saga)
	case message.Type_TYPE_EVENT_CARD_AUTHORIZED:
		strategy = event_handler.NewNextStepSagaWhenCardAuthorizedHandler(saga)
	case message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:
		strategy = event_handler.NewNextStepSagaWhenCardAuthorizeFailedHandler(saga)

	default:
		return nil, errors.New("not implemented event by CreateOrderSagaContext")
	}

	return &CreateOrderSagaContext{strategy: strategy, mes: mes}, nil
}

func (c *CreateOrderSagaContext) Handler(ctx context.Context) error {
	return c.strategy.Handler(ctx, c.mes)
}
