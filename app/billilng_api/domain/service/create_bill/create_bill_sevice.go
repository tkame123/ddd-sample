package create_ticket

import (
	"context"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/model"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/service"
	"log"
)

type CreatBillService struct {
}

func (c CreatBillService) AuthorizeCard(ctx context.Context, orderID model.OrderID, token any) error {
	//TODO implement me
	log.Println("implement me: AuthorizeCard")
	return nil
}

func NewService() service.CreateBill {
	return &CreatBillService{}
}
