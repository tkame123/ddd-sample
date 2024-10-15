package model

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type OrderID = uuid.UUID

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
