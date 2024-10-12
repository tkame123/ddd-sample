package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenOrderCreatedHandler struct{}

func NewNextStepSagaWhenOrderCreatedHandler() domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenOrderCreatedHandler{}
}

func (h *NextStepSagaWhenOrderCreatedHandler) Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_ORDER_CREATED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_CreteTicket); err != nil {
		return err
	}

	return nil
}
