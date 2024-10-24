package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type callOrderCancelSagaWhenCardCanceledHandler struct {
	rep        repository.Repository
	orderSVC   service.OrderService
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewCallOrderCancelSagaWhenCardCanceledHandler(
	rep repository.Repository,
	orderSVC service.OrderService,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) domain_event.EventHandler {
	return &callOrderCancelSagaWhenCardCanceledHandler{
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (h *callOrderCancelSagaWhenCardCanceledHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_CARD_CANCELED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.EventCardCanceled
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}
	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	state, err := h.rep.CancelOrderSagaStateFindOne(ctx, *id)
	if err != nil {
		return err
	}
	saga, err := cancel_order_saga.NewCancelOrderSaga(
		state,
		h.orderSVC,
		h.kitchenAPI,
		h.billingAPI,
	)
	if err != nil {
		return err
	}

	err = saga.Event(ctx, mes)
	if err != nil {
		return err
	}

	newState := saga.ExportState()
	err = h.rep.CancelOrderSagaStateSave(ctx, newState)
	if err != nil {
		return err
	}

	return nil
}
