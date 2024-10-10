package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenTicketCreationFailedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketCreationFailedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenTicketCreationFailedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketCreationFailedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_TicketCreationFailed {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_TicketCreationFailed); err != nil {
		return err
	}

	return nil
}
