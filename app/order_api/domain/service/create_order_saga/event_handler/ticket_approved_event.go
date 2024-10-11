package event_handler

import (
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenTicketApprovedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenTicketApprovedHandler(saga *servive.CreateOrderSaga) domain_event.EventHandler {
	return &NextStepSagaWhenTicketApprovedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenTicketApprovedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_TicketApproved {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveOrder); err != nil {
		return err
	}

	return nil
}
