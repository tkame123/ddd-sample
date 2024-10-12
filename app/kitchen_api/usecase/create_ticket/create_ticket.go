package create_ticket

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

func (s *CreatTicketService) CreateTicket(ctx context.Context, orderID model.OrderID, items []*model.TicketItemRequest) error {
	ticket, events, err := model.NewTicket(orderID, items)
	if err != nil {
		return err
	}

	if err := s.rep.TicketSave(ctx, ticket); err != nil {
		return err
	}

	s.pub.PublishMessages(ctx, events)

	return nil
}
