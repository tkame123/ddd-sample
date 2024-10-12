package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketApprovedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketApprovedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenTicketApprovedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketApprovedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_APPROVED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveOrder); err != nil {
		return err
	}

	return nil
}
