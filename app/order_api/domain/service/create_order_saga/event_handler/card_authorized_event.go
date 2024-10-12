package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenCardAuthorizedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenCardAuthorizedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenCardAuthorizedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenCardAuthorizedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_CARD_AUTHORIZED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveTicket); err != nil {
		return err
	}

	return nil
}
