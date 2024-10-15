package proxy

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
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

func (k *KitchenAPI) CreateTicket(ctx context.Context, orderID model.OrderID) error {
	order, err := k.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to find order: %w", err)
	}

	items := make([]*message.CommandTicketCreate_Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &message.CommandTicketCreate_Item{
			ItemId:   item.ItemID.String(),
			Quantity: int32(item.Quantity),
		})
	}

	command, err := model.CreateMessage(
		&message.CommandTicketCreate{
			OrderId: order.OrderID.String(),
			Items:   items,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	k.pub.PublishMessages(ctx, []*message.Message{command})

	return nil
}

func (k *KitchenAPI) ApproveTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
	//	TODO: Implement this
	log.Println("KitchenAPI ApproveTicket")
	return nil
}

func (k *KitchenAPI) RejectTicket(ctx context.Context, orderID model.OrderID, ticketID model.TicketID) error {
	//	TODO: Implement this
	log.Println("KitchenAPI RejectTicket")
	return nil
}
