package create_order_saga

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type eventHandler struct {
	rep        repository.Repository
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewEventHandler(
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) domain_event.EventHandler {
	return &eventHandler{
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (h *eventHandler) Handler(ctx context.Context, m *message.Message) error {
	var id *model.OrderID

	switch m.Subject.Type {
	case message.Type_TYPE_EVENT_ORDER_CREATED:
		var v message.EventOrderCreated
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_ORDER_APPROVED:
		var v message.EventOrderApproved
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_ORDER_REJECTED:
		var v message.EventOrderRejected
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_TICKET_CREATED:
		var v message.EventTicketCreated
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:
		var v message.EventTicketCreationFailed
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_TICKET_APPROVED:
		var v message.EventTicketApproved
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_TICKET_REJECTED:
		var v message.EventTicketRejected
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_CARD_AUTHORIZED:
		var v message.EventCardAuthorized
		err := m.Envelope.UnmarshalTo(&v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		i, err := model.OrderIdParse(v.OrderId)
		if err != nil {
			return fmt.Errorf("failed to parse order id: %w", err)
		}
		id = i
	case message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:
		var v message.EventCardAuthorizationFailed
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

	state, err := h.rep.CreateOrderSagaStateFindOne(ctx, *id)
	if err != nil {
		return err
	}
	saga, err := NewCreateOrderSaga(
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

	return nil
}
