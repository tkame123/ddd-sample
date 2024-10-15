package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenTicketCreatedHandler struct {
	rep repository.Repository
}

func NewNextStepSagaWhenTicketCreatedHandler(rep repository.Repository) domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenTicketCreatedHandler{rep: rep}
}

func (h *NextStepSagaWhenTicketCreatedHandler) Handler(ctx context.Context, sagaFactory domain_event.SagaFactory, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_TICKET_CREATED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.EventTicketCreated
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

	if err := saga.Event(ctx, mes); err != nil {
		return err
	}

	return nil
}
