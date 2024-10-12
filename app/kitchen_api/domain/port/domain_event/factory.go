package domain_event

import "github.com/tkame123/ddd-sample/lib/event_helper"

type EventFactory interface {
	Event() (event_helper.Event, error)
}
