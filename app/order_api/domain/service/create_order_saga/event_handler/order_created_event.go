package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenOrderCreatedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenOrderCreatedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenOrderCreatedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenOrderCreatedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_ORDER_CREATED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_CreteTicket); err != nil {
		return err
	}

	return nil
}
