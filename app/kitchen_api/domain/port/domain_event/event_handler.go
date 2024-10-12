package domain_event

import (
	"context"
	ev "github.com/tkame123/ddd-sample/lib/event_helper"
)

type EventHandler interface {
	Handler(ctx context.Context, event ev.Event) error
}
