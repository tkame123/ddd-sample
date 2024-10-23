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

type cancelOrderSagaEventHandler struct {
	rep        repository.Repository
	orderSVC   service.CancelOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewCreateCancelSagaEventHandler(
	rep repository.Repository,
	orderSVC service.CancelOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) domain_event.EventHandler {
	return &cancelOrderSagaEventHandler{
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (h *cancelOrderSagaEventHandler) Handler(ctx context.Context, m *message.Message) error {
	var id *model.OrderID

	switch m.Subject.Type {
	case message.Type_TYPE_EVENT_ORDER_CANCELED:
		var v message.EventOrderCanceled
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	case message.Type_TYPE_EVENT_TICKET_CANCELED:
		var v message.EventTicketCanceled
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	case message.Type_TYPE_EVENT_CARD_CANCELED:
		var v message.EventCardCanceled
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	case message.Type_TYPE_EVENT_ORDER_CANCELLATION_CONFIRMED:
		var v message.EventOrderCancellationConfirmed
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	case message.Type_TYPE_EVENT_TICKET_CANCELLATION_REJECTED:
		var v message.EventTicketCancellationRejected
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	case message.Type_TYPE_EVENT_ORDER_CANCELLATION_REJECTED:
		var v message.EventOrderCancellationRejected
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i

	default:
		return fmt.Errorf("invalid event type: %s", m.Subject.Type)
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

	err = saga.Event(ctx, m)
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
