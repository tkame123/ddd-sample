package model

import "github.com/tkame123/ddd-sample/app/delivery_api/domain/event"

// 集約ルート
type Shipment struct {
	shipmentID    ShipmentID
	orderID       OrderID
	shipmentItems []*ShipmentItem
}

type ShipmentItem struct {
	ShipmentID ShipmentID
	ItemID     ItemID
	Quantity   int64
}

type ShipmentItemRequest struct {
	ItemID   ItemID
	quantity int64
}

func NewShipment(orderID OrderID, items []*ShipmentItemRequest) (*Shipment, []event.ShipmentEvent, error) {
	shipmentID := generateID()

	shipmentItems := make([]*ShipmentItem, 0, len(items))
	for _, item := range items {
		shipmentItems = append(shipmentItems, &ShipmentItem{
			ShipmentID: shipmentID,
			ItemID:     item.ItemID,
			Quantity:   item.quantity,
		})
	}

	shipment := &Shipment{
		shipmentID:    shipmentID,
		orderID:       orderID,
		shipmentItems: shipmentItems,
	}

	createEvent := event.NewShipmentCreated(shipment.ShipmentID())

	return shipment, []event.ShipmentEvent{createEvent}, nil
}

func (s *Shipment) ShipmentID() ShipmentID {
	return s.shipmentID
}
