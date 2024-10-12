package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketCreationFailedHandler struct{}

func NewNextStepSagaWhenTicketCreationFailedHandler() domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenTicketCreationFailedHandler{}
}

func (h *NextStepSagaWhenTicketCreationFailedHandler) Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_CREATION_FAILED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_TicketCreationFailed); err != nil {
		return err
	}

	return nil
}
