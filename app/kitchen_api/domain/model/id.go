package model

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/lib/event"
)

type TicketID = uuid.UUID
type OrderID = uuid.UUID
type ItemID = uuid.UUID

func generateID() uuid.UUID {
	id := uuid.New()

	return id
}

func TicketIdParse(id string) (*TicketID, error) {
	parsedId, err := event.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}

func OrderIdParse(id string) (*OrderID, error) {
	parsedId, err := event.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}

func ItemIdParse(id string) (*ItemID, error) {
	parsedId, err := event.ParseID(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}
