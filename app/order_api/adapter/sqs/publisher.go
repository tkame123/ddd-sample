package sqs

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
)

type eventPublisher struct {
}

func NewEventPublisher() domain_event.Publisher {
	return &eventPublisher{}
}

func (s *eventPublisher) PublishMessages(ctx context.Context, events []model.OrderEvent) {
	panic("implement me")
}
