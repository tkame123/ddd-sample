package proxy

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/lib/event_helper"
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

	command := &TicketCreateCommand{
		OrderID: orderID,
	}

	if len(order.OrderItems) > 0 {
		items := make([]*TicketItemRequest, 0, len(order.OrderItems))
		for _, item := range order.OrderItems {
			items = append(items, &TicketItemRequest{
				ItemID:   item.ItemID,
				Quantity: item.Quantity,
			})
		}
		command.Items = items
	}

	k.pub.PublishMessages(ctx, []event_helper.Event{command})
}

func (k *KitchenAPI) ApproveTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI ApproveTicket")
}

func (k *KitchenAPI) RejectTicket(ctx context.Context, orderID model.OrderID) {
	//	TODO: Implement this
	log.Println("KitchenAPI RejectTicket")
}

// TODO: eventの根本的なリファクタやるまえに一旦使いたいのでここにおく。。
type TicketItemRequest struct {
	ItemID   model.ItemID `json:"item_id"`
	Quantity int          `json:"quantity"`
}

type TicketCreateCommand struct {
	OrderID model.OrderID        `json:"order_id"`
	Items   []*TicketItemRequest `json:"items"`
}

func (e *TicketCreateCommand) Name() string {
	return event_helper.CommandName_TicketCreate
}

func (e *TicketCreateCommand) ID() uuid.UUID {
	return uuid.Nil
}

func (e *TicketCreateCommand) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
