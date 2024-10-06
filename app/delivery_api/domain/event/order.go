package event

import (
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/model"
	"github.com/tkame123/ddd-sample/lib/event"
)

type ShipmentEvent interface {
	event.Event
	ID() model.OrderID
}

type ShipmentCreated struct {
	shipmentID model.ShipmentID
}

func NewShipmentCreated(shipmentID model.ShipmentID) *ShipmentCreated {
	return &ShipmentCreated{
		shipmentID: shipmentID,
	}
}

func (e *ShipmentCreated) Name() string {
	return "event.delivery.created"
}

func (e *ShipmentCreated) ID() model.ShipmentID {
	return e.shipmentID
}

type OrderItemsUpdated struct {
	orderID model.OrderID
}
