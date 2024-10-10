package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenCardAuthorizedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenCardAuthorizedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenCardAuthorizedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenCardAuthorizedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_CardAuthorized {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveTicket); err != nil {
		return err
	}

	return nil
}
