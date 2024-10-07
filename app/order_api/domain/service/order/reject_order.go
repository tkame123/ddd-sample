package order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type RejectOrderInput struct {
	OrderID model.OrderID
}

type RejectOrderOutput struct {
	OrderID model.OrderID
}

func (s *Service) RejectOrder(ctx context.Context, input RejectOrderInput) (*RejectOrderOutput, error) {
	order, err := s.rep.Order.FindOne(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	events, err := order.RejectOrder()
	if err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &RejectOrderOutput{OrderID: order.OrderID()}, nil
}
