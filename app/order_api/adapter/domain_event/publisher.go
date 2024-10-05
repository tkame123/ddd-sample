package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/event"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []event.OrderEvent)
}
