package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/model"
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/port/repository"
)

type CreateShipment struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewCreateShipment(rep repository.Repository) *CreateShipment {
	return &CreateShipment{
		rep: rep,
	}
}

type CreateShipmentInput struct {
	OrderID model.OrderID
	Items   []*model.ShipmentItemRequest
}

type CreateShipmentOutput struct {
	ShipmentID model.ShipmentID
}

func (c *CreateShipment) Execute(ctx context.Context, input CreateShipmentInput) (*CreateShipmentOutput, error) {
	shipment, events, err := model.NewShipment(input.OrderID, input.Items)
	if err != nil {
		return nil, err
	}

	if err := c.rep.Shipment.Save(ctx, shipment); err != nil {
		return nil, err
	}

	c.pub.PublishMessages(ctx, events)

	return &CreateShipmentOutput{ShipmentID: shipment.ShipmentID()}, nil
}
