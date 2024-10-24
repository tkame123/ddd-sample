package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type ticketService struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewTicketService(rep repository.Repository, pub domain_event.Publisher) service.Ticket {
	return &ticketService{
		rep: rep,
		pub: pub,
	}
}

func (s *ticketService) CreateTicket(ctx context.Context, orderID model.OrderID, items []*model.TicketItemRequest) error {
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

func (s *ticketService) ApproveTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
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

func (s *ticketService) RejectTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
	ticket, err := s.rep.TicketFindOneByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	events, err := ticket.RejectTicket()
	if err != nil {
		return err
	}

	s.pub.PublishMessages(ctx, events)

	return nil
}

func (c *ticketService) CancelTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
	// TODO: Repository実装後に仕上げる必要有り

	mes, err := model.CreateMessage(
		&message.EventTicketCanceled{
			OrderId:  orderID.String(),
			TicketId: ticketID.String(),
		},
	)
	if err != nil {
		return err
	}

	c.pub.PublishMessages(ctx, []*message.Message{mes})

	return nil
}
