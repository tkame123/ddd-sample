package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type NextStepSagaWhenOrderApprovedHandler struct{}

func NewNextStepSagaWhenOrderApprovedHandler() domain_event.CreateOrderSagaEventHandler {
	return &NextStepSagaWhenOrderApprovedHandler{}
}

func (h *NextStepSagaWhenOrderApprovedHandler) Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_EVENT_ORDER_APPROVED {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	if err := saga.Event(ctx, servive.CreateOrderSagaEvent_OrderApprove); err != nil {
		return err
	}

	return nil
}
