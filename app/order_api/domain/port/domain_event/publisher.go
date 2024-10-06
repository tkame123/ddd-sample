package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []model.OrderEvent)
}
