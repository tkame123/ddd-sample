package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenCardAuthorizedHandler struct {
	rep repository.Repository
}

func NewNextStepSagaWhenCardAuthorizedHandler(rep repository.Repository) domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenCardAuthorizedHandler{rep: rep}
}

func (h *NextStepSagaWhenCardAuthorizedHandler) Handler(ctx context.Context, sagaFactory domain_event.SagaFactory, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_CARD_AUTHORIZED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.EventCardAuthorized
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}
	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	saga, err := sagaFactory(ctx, h.rep, *id)
	if err != nil {
		return err
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_ApproveTicket); err != nil {
		return err
	}

	return nil
}
