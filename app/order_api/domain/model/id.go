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
