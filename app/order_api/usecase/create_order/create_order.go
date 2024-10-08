package create_order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (s *s) CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error) {
	order, events, err := model.NewOrder(items)
	if err != nil {
		return "", err
	}

	if err := s.rep.Order.Save(ctx, order); err != nil {
		return "", err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID(), nil
}
