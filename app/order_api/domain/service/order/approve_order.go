package order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type ApproveOrderInput struct {
	OrderID model.OrderID
}

type ApproveOrderOutput struct {
	OrderID model.OrderID
}

func (s *Service) ApproveOrder(ctx context.Context, input ApproveOrderInput) (*ApproveOrderOutput, error) {
	order, err := s.rep.Order.FindOne(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	events, err := order.ApproveOrder()
	if err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &ApproveOrderOutput{OrderID: order.OrderID()}, nil
}
