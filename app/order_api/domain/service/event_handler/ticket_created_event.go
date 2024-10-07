package event_handler

import (
	"context"
	"errors"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenTicketCreatedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketCreatedHandler(saga *servive.CreateOrderSaga) *NextStepSagaWhenTicketCreatedHandler {
	return &NextStepSagaWhenTicketCreatedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketCreatedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_TicketCreated {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_RejectOrder); err != nil {
		return err
	}

	return nil
}
