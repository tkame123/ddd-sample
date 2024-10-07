package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenOrderCreatedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenOrderCreatedHandler(saga *servive.CreateOrderSaga) *NextStepSagaWhenOrderCreatedHandler {
	return &NextStepSagaWhenOrderCreatedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenOrderCreatedHandler) Handler(ctx context.Context, event model.OrderCreatedEvent) error {
	if event.Name() != ev.EventName_OrderCreated {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_CreteTicket); err != nil {
		return err
	}

	return nil
}
