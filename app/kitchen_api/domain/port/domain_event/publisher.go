package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/lib/event"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []event.Event)
}
