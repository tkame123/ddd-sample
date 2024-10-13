package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"log"
)

func (r repo) TicketFindOne(ctx context.Context, id model.TicketID) (*model.Ticket, error) {
	//TODO implement me
	log.Println("implement me: TicketFindOne")
	return nil, nil
}

func (r repo) TicketFindOneByOrderID(ctx context.Context, id model.OrderID) (*model.Ticket, error) {
	//TODO implement me
	log.Println("implement me: TicketFindOneByOrderID")
	return nil, nil
}

func (r repo) TicketSave(ctx context.Context, order *model.Ticket) error {
	//TODO implement me
	log.Println("implement me: TicketSave")
	return nil
}
