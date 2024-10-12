package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketRejectedHandler struct {
}

func NewNextStepSagaWhenTicketRejectedHandler() domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenTicketRejectedHandler{}
}

func (h *NextStepSagaWhenTicketRejectedHandler) Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_REJECTED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_RejectOrder); err != nil {
		return err
	}

	return nil
}
