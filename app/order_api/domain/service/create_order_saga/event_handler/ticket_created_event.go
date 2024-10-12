package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketCreatedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketCreatedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenTicketCreatedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketCreatedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_CREATED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_AuthorizeCard); err != nil {
		return err
	}

	return nil
}
