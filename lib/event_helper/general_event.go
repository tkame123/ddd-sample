package event_helper

import (
	"encoding/json"
	"github.com/google/uuid"
)

// Deprecated: use ProtoEvent
type GeneralEvent struct {
	Id     uuid.UUID
	Type   Name
	Origin json.RawMessage
}

// Deprecated: use ProtoEvent
func ParseID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

// Deprecated: use ProtoEvent
func NewGeneralEvent(id uuid.UUID, name Name) Event {
	return &GeneralEvent{
		Id:   id,
		Type: name,
	}
}

// Deprecated: use ProtoEvent
func NewGeneralEventFromRaw(raw RawEvent) (Event, error) {
	id, err := uuid.Parse(raw.ID)
	if err != nil {
		return nil, err
	}

	return &GeneralEvent{
		Id:     id,
		Type:   raw.Type,
		Origin: raw.Origin,
	}, nil
}

// Deprecated: use ProtoEvent
func (e *GeneralEvent) Name() Name {
	return e.Type
}

// Deprecated: use ProtoEvent
func (e *GeneralEvent) ID() uuid.UUID {
	return e.Id
}

// Deprecated: use ProtoEvent
func (e *GeneralEvent) ToBody() (string, error) {
	return string(e.Origin), nil
}
