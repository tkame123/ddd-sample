package create_ticket

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/proto/message"
)

func (s *CreatTicketService) ApproveTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
	mes, err := model.CreateMessage(
		&message.EventTicketApproved{
			OrderId:  orderID.String(),
			TicketId: ticketID.String(),
		},
	)
	if err != nil {
		return err
	}

	s.pub.PublishMessages(ctx, []*message.Message{mes})

	return nil

	// TODO: repositoryの実装待ち
	//ticket, err := s.rep.TicketFindOneByOrderID(ctx, orderID)
	//if err != nil {
	//	return err
	//}
	//
	//events, err := ticket.ApproveTicket()
	//if err != nil {
	//	return err
	//}
	//
	//s.pub.PublishMessages(ctx, events)
	//
	//return nil
}
