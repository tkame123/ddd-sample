package create_order

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (s *s) CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error) {
	order, events, err := model.NewOrder(items)
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.CreateOrderSagaStateSave(ctx, &model.CreateOrderSagaState{
		OrderID: order.OrderID,
		Current: model.CreateOrderSagaStep_ApprovalPending,
	}); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}
