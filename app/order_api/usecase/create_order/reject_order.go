package create_order

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (s *s) RejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return "", err
	}

	events, err := order.RejectOrder()
	if err != nil {
		return "", err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID(), nil
}
