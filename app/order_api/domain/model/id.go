package model

import (
	"github.com/google/uuid"
)

type OrderID = uuid.UUID
type OrderItemID = uuid.UUID
type ItemID = uuid.UUID

func generateID() uuid.UUID {
	id := uuid.New()

	return id
}

func OrderIdParse(id string) (*OrderID, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &parsedId, nil
}
