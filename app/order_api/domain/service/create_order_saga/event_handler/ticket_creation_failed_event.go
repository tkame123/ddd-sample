package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketCreationFailedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketCreationFailedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenTicketCreationFailedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketCreationFailedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_CREATION_FAILED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_TicketCreationFailed); err != nil {
		return err
	}

	return nil
}
