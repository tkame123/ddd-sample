package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/model"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type CreatBillService struct {
	pub domain_event.Publisher
}

func NewService(pub domain_event.Publisher) service.CreateBill {
	return &CreatBillService{pub: pub}
}

func (c CreatBillService) AuthorizeCard(ctx context.Context, orderID model.OrderID, token any) error {
	//TODO implement me
	log.Println("implement me: AuthorizeCard")

	// TODO: 仮置き
	m, err := model.CreateMessage(
		&message.EventCardAuthorized{
			OrderId: orderID.String(),
		})
	if err != nil {
		return err
	}

	c.pub.PublishMessages(ctx, []*message.Message{m})

	return nil
}
