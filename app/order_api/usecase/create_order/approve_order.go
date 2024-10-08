package create_order

import (
	"context"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
)

func (s *s) ApproveOrder(ctx context.Context, input servive.ApproveOrderInput) (*servive.ApproveOrderOutput, error) {
	order, err := s.rep.Order.FindOne(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	events, err := order.ApproveOrder()
	if err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &servive.ApproveOrderOutput{OrderID: order.OrderID()}, nil
}
