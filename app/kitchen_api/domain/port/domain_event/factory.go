package domain_event

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/proto/message"
)

type Message interface {
	ID() uuid.UUID
	Raw() *message.Message
}

type MessageFactory interface {
	Event() (Message, error)
}
