package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenOrderRejectedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenOrderRejectedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenOrderRejectedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenOrderRejectedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_OrderRejected {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_RejectedOrder); err != nil {
		return err
	}

	return nil
}
