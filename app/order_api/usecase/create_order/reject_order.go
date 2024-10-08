package create_order

import (
	"context"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
)

func (s *s) RejectOrder(ctx context.Context, input servive.RejectOrderInput) (*servive.RejectOrderOutput, error) {
	order, err := s.rep.Order.FindOne(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	events, err := order.RejectOrder()
	if err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &servive.RejectOrderOutput{OrderID: order.OrderID()}, nil
}
