package event_handler

import (
	"context"
	"errors"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	ev "github.com/tkame123/ddd-sample/lib/event"
)

type NextStepSagaWhenCardAuthorizeFailedHandler struct {
	saga *servive.CreateOrderSaga
}

func NewNextStepSagaWhenCardAuthorizeFailedHandler(saga *servive.CreateOrderSaga) *NextStepSagaWhenCardAuthorizeFailedHandler {
	return &NextStepSagaWhenCardAuthorizeFailedHandler{
		saga: saga,
	}
}

func (h *NextStepSagaWhenCardAuthorizeFailedHandler) Handler(ctx context.Context, event ev.Event) error {
	if event.Name() != ev.EventName_CardAuthorizeFailed {
		return errors.New("invalid event")
	}

	if err := h.saga.Event(ctx, servive.CreateOrderSagaEvent_AuthorizeCardFailed); err != nil {
		return err
	}

	return nil
}
