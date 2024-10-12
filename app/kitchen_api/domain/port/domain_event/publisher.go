package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/lib/event_helper"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []event_helper.Event)
}
