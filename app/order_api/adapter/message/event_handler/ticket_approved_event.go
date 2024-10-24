package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type callOrderCreateSagaWhenTicketApprovedHandler struct {
	rep        repository.Repository
	orderSVC   service.OrderService
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewCallOrderCreateSagaWhenTicketApprovedHandler(
	rep repository.Repository,
	orderSVC service.OrderService,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) domain_event.EventHandler {
	return &callOrderCreateSagaWhenTicketApprovedHandler{
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (h *callOrderCreateSagaWhenTicketApprovedHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_APPROVED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.EventTicketApproved
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}
	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	state, err := h.rep.CreateOrderSagaStateFindOne(ctx, *id)
	if err != nil {
		return err
	}
	saga, err := create_order_saga.NewCreateOrderSaga(
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
	err = h.rep.CreateOrderSagaStateSave(ctx, newState)
	if err != nil {
		return err
	}

	return nil
}
