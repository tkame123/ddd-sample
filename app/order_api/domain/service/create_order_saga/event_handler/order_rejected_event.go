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

type NextStepSagaWhenOrderRejectedHandler struct {
	rep repository.Repository
}

func NewNextStepSagaWhenOrderRejectedHandler(rep repository.Repository) domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenOrderRejectedHandler{rep: rep}
}

func (h *NextStepSagaWhenOrderRejectedHandler) Handler(ctx context.Context, sagaFactory domain_event.SagaFactory, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_ORDER_REJECTED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.EventOrderRejected
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

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_RejectedOrder); err != nil {
		return err
	}

	return nil
}
