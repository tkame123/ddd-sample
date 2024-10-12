package domain_event

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/proto/message"
)

type SagaMessage interface {
	ID() uuid.UUID
	Raw() *message.Message
}

type SagaEventFactory interface {
	Event() (SagaMessage, error)
}
