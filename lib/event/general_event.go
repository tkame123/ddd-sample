package event

import (
	"encoding/json"
	"github.com/google/uuid"
)

type GeneralEvent struct {
	Id     uuid.UUID
	Type   Name
	Origin json.RawMessage
}

func ParseID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func NewGeneralEvent(id uuid.UUID, name Name) Event {
	return &GeneralEvent{
		Id:   id,
		Type: name,
	}
}

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

func (e *GeneralEvent) Name() Name {
	return e.Type
}

func (e *GeneralEvent) ID() uuid.UUID {
	return e.Id
}

func (e *GeneralEvent) ToBody() (string, error) {
	return string(e.Origin), nil
}
