package fake

import (
	"github.com/google/uuid"
	event2 "github.com/tkame123/ddd-sample/lib/event"
)

type GeneralEvent struct {
	id   uuid.UUID
	name string
}

func NewGeneralEvent(id uuid.UUID, name string) *GeneralEvent {
	return &GeneralEvent{
		id:   id,
		name: name,
	}
}

func (e *GeneralEvent) Name() event2.Name {
	return e.name
}

func (e *GeneralEvent) ID() uuid.UUID {
	return e.id
}
