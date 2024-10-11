package proxy

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"log"
)

type KitchenAPI struct {
}

func NewKitchenAPI() external_service.KitchenAPI {
	return &KitchenAPI{}
}

func (k *KitchenAPI) CreateTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI CreateTicket")
}

func (k *KitchenAPI) ApproveTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI ApproveTicket")
}

func (k *KitchenAPI) RejectTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI RejectTicket")
}
