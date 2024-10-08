package create_order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
)

func (s *s) CreateOrder(ctx context.Context, input servive.CreateOrderInput) (*servive.CreateOrderOutput, error) {
	order, events, err := model.NewOrder(input.Items)
	if err != nil {
		return nil, err
	}

	if err := s.rep.Order.Save(ctx, order); err != nil {
		return nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return &servive.CreateOrderOutput{OrderID: order.OrderID()}, nil
}
