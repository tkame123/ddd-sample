package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenOrderApprovedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenOrderApprovedHandler(saga *servive.CreateOrderSaga) *NextStepSagaWhenOrderApprovedHandler {
	return &NextStepSagaWhenOrderApprovedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenOrderApprovedHandler) Handler(ctx context.Context, event model.OrderApprovedEvent) error {
	if event.Name() != ev.EventName_OrderApproved {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_OrderApprove); err != nil {
		return err
	}

	return nil
}