package domain_event

import (
	"github.com/tkame123/ddd-sample/proto/message"
)

type EventFactory interface {
	Event() (*message.Message, error)
}
