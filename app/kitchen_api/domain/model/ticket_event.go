package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type TicketCreateCommand struct {
	OrderID OrderID              `json:"order_id"`
	Items   []*TicketItemRequest `json:"items"`
}

func (e *TicketCreateCommand) Name() string {
	return event_helper.CommandName_TicketCreate
}

func (e *TicketCreateCommand) ID() TicketID {
	return uuid.Nil
}

func (e *TicketCreateCommand) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type TicketCreatedEvent struct {
	TicketID TicketID
}

func (e *TicketCreatedEvent) Name() string {
	return event_helper.EventName_TicketCreated
}

func (e *TicketCreatedEvent) ID() TicketID {
	return e.TicketID
}

func (e *TicketCreatedEvent) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type TicketCreationFailedEvent struct {
	TicketID TicketID
}

func (e *TicketCreationFailedEvent) Name() string {
	return event_helper.EventName_TicketCreationFailed
}

func (e *TicketCreationFailedEvent) ID() TicketID {
	return e.TicketID
}

func (e *TicketCreationFailedEvent) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type TicketApprovedEvent struct {
	TicketID TicketID
}

func (e *TicketApprovedEvent) Name() string {
	return event_helper.EventName_TicketApproved
}

func (e *TicketApprovedEvent) ID() TicketID {
	return e.TicketID
}

func (e *TicketApprovedEvent) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type TicketRejectedEvent struct {
	TicketID TicketID
}

func (e *TicketRejectedEvent) Name() string {
	return event_helper.EventName_TicketRejected
}

func (e *TicketRejectedEvent) ID() TicketID {
	return e.TicketID
}

func (e *TicketRejectedEvent) ToBody() (string, error) {
	var raw event_helper.RawEvent
	raw.Type = e.Name()
	raw.ID = e.ID().String()
	originByte, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	raw.Origin = originByte
	body, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
