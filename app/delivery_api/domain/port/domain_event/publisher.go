package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/event"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []event.ShipmentEvent)
}
