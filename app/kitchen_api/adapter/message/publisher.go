package message

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"log"
)

type EventPublisher struct {
	publisher *sns.Publisher
}

func NewEventPublisher(publisher *sns.Publisher) domain_event.Publisher {
	return &EventPublisher{publisher: publisher}
}

func (s *EventPublisher) PublishMessages(ctx context.Context, events []event_helper.Event) {
	for _, e := range events {
		body, err := e.ToBody()
		if err != nil {
			log.Printf("failed to marshal event %v", err)
			continue
		}
		if err := s.publisher.PublishMessage(ctx, e.Name(), body); err != nil {
			// TODO Transactional Outbox Patternの導入までは通知して手動対応ってなるのだろうか。。。
			log.Printf("failed to publish message %v", err)
		}
	}
}
