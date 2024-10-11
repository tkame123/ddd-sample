package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenOrderApprovedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenOrderApprovedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenOrderApprovedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenOrderApprovedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_OrderApproved {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_OrderApprove); err != nil {
		return err
	}

	return nil
}
