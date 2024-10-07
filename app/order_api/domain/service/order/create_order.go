package order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrderInput struct {
	Items []*model.OrderItemRequest
}

type CreateOrderOutput struct {
	OrderID model.OrderID
}

func (s *Service) CreateOrder(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error) {
	order, events, err := model.NewOrder(input.Items)
	if err != nil {
		return nil, err
	}

	if err := s.rep.Order.Save(ctx, order); err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &CreateOrderOutput{OrderID: order.OrderID()}, nil
}
