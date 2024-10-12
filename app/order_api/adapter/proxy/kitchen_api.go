package proxy

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type KitchenAPI struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewKitchenAPI(
	rep repository.Repository,
	pub domain_event.Publisher,
) external_service.KitchenAPI {
	return &KitchenAPI{rep: rep, pub: pub}
}

func (k *KitchenAPI) CreateTicket(ctx context.Context, orderID model.OrderID) {
	order, err := k.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		log.Printf("failed to find order: %v", err)
		return
	}

	items := make([]*message.CommandTicketCreate_Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &message.CommandTicketCreate_Item{
			ItemId:   item.ItemID.String(),
			Quantity: int32(item.Quantity),
		})
	}

	command, err := event_helper.CreateMessage(
		message.Type_TYPE_COMMAND_TICKET_CREATE,
		message.Service_SERVICE_ORDER,
		&message.CommandTicketCreate{
			OrderId: order.OrderID.String(),
			Items:   items,
		},
	)

	k.pub.PublishMessages(ctx, []*message.Message{command})
}

func (k *KitchenAPI) ApproveTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI ApproveTicket")
}

func (k *KitchenAPI) RejectTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI RejectTicket")
}
