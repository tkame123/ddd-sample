package model

import "github.com/google/uuid"

type OrderID = string
type OrderItemID = string
type ItemID = string

func generateID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id.String()
}
