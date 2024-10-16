package create_order

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (s *s) RejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.RejectOrder()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}
