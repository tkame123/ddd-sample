package model

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type OrderID = uuid.UUID
type OrderItemID = uuid.UUID
type ItemID = uuid.UUID

type TicketID = uuid.NullUUID

func generateID() uuid.UUID {
	id := uuid.New()

	return id
}

func OrderIdParse(id string) (*OrderID, error) {
	parsedId, err := event_helper.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}

func OrderItemIdParse(id string) (*OrderItemID, error) {
	parsedId, err := event_helper.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}

func ItemIdParse(id string) (*ItemID, error) {
	parsedId, err := event_helper.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}

func TicketIdParse(id string) (*TicketID, error) {
	parsedId, err := event_helper.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &uuid.NullUUID{Valid: true, UUID: parsedId}, nil
}
