package create_ticket

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

func (s *CreatTicketService) ApproveTicket(ctx context.Context, orderID model.OrderID) error {
	ticket, err := s.rep.TicketFindOneByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	events, err := ticket.ApproveTicket()
	if err != nil {
		return err
	}

	s.pub.PublishMessages(ctx, events)

	return nil
}
