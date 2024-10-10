package model

import (
	"encoding/json"
	"github.com/tkame123/ddd-sample/lib/event"
)

type OrderCreatedEvent struct {
	OrderID OrderID
}

func (e *OrderCreatedEvent) Name() event.Name {
	return event.EventName_OrderCreated
}

func (e *OrderCreatedEvent) ID() OrderID {
	return e.OrderID
}

func (e *OrderCreatedEvent) ToBody() (string, error) {
	var dto struct {
		Type   string          `json:"type"`
		Origin json.RawMessage `json:"origin"`
	}
	dto.Type = e.Name()
	orignByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	dto.Origin = orignByte
	body, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type OrderApprovedEvent struct {
	OrderID OrderID
}

func (e *OrderApprovedEvent) Name() event.Name {
	return event.EventName_OrderApproved
}

func (e *OrderApprovedEvent) ID() OrderID {
	return e.OrderID
}

func (e *OrderApprovedEvent) ToBody() (string, error) {
	var dto struct {
		Type   string          `json:"type"`
		Origin json.RawMessage `json:"origin"`
	}
	dto.Type = e.Name()
	orignByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	dto.Origin = orignByte
	body, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type OrderRejectedEvent struct {
	OrderID OrderID
}

func (e *OrderRejectedEvent) Name() event.Name {
	return event.EventName_OrderRejected
}

func (e *OrderRejectedEvent) ID() OrderID {
	return e.OrderID
}

func (e *OrderRejectedEvent) ToBody() (string, error) {
	var dto struct {
		Type   string          `json:"type"`
		Origin json.RawMessage `json:"origin"`
	}
	dto.Type = e.Name()
	orignByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	dto.Origin = orignByte
	body, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
