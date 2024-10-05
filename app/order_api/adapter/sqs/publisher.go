package sqs

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/event"
)

type eventPublisher struct {
}

func NewEventPublisher() domain_event.Publisher {
	return &eventPublisher{}
}

func (s *eventPublisher) PublishMessages(ctx context.Context, events []event.OrderEvent) {
	panic("implement me")
}
