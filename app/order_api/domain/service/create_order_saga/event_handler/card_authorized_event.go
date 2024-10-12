package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenCardAuthorizedHandler struct{}

func NewNextStepSagaWhenCardAuthorizedHandler() domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenCardAuthorizedHandler{}
}

func (h *NextStepSagaWhenCardAuthorizedHandler) Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_CARD_AUTHORIZED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveTicket); err != nil {
		return err
	}

	return nil
}
