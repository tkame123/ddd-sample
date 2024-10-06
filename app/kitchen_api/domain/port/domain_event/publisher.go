package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []model.TicketEvent)
}
