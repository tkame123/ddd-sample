package domain_event

import "github.com/tkame123/ddd-sample/lib/event"

type EventFactory interface {
	Event() (event.Event, error)
}
