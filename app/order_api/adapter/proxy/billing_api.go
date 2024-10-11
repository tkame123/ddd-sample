package proxy

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"log"
)

type BillingAPI struct {
}

func NewBillingAPI() external_service.BillingAPI {
	return &BillingAPI{}
}

func (k *BillingAPI) AuthorizeCard(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("BillingAPI AuthorizeCard")
}
