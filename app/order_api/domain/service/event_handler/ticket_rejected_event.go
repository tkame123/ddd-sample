package event_handler

import (
	"context"
	"errors"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenTicketRejectedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketRejectedHandler(saga *servive.CreateOrderSaga) *NextStepSagaWhenTicketRejectedHandler {
	return &NextStepSagaWhenTicketRejectedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketRejectedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_TicketRejected {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_RejectOrder); err != nil {
		return err
	}

	return nil
}
