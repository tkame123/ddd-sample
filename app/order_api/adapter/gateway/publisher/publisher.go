package publisher

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event"
	"log"
)

type topicArn = string

type EventPublisher struct {
	sns      sns.Publisher
	topicMap map[event.Name]topicArn
}

func NewEventPublisher(envCfg *provider.EnvConfig, sns *sns.Publisher) domain_event.Publisher {
	topicMap := map[event.Name]topicArn{
		event.EventName_OrderCreated:    envCfg.ArnTopicEventOrderOrderCreated,
		event.EventName_OrderApproved:   envCfg.ArnTopicEventOrderOrderApproved,
		event.EventName_OrderRejected:   envCfg.ArnTopicEventOrderOrderRejected,
		event.CommandName_TicketCreate:  envCfg.ArnTopicCommandKitchenTicketCreate,
		event.CommandName_TicketApprove: envCfg.ArnTopicCommandOrderOrderApprove,
		event.CommandName_TicketReject:  envCfg.ArnTopicCommandOrderOrderReject,
		event.CommandName_CardAuthorize: envCfg.ArnTopicCommandBillingCardAuthorize,
	}

	return &EventPublisher{sns: *sns, topicMap: topicMap}
}

func (s *EventPublisher) PublishMessages(ctx context.Context, events []event.Event) {
	for _, e := range events {
		var topic string
		topic, ok := s.topicMap[e.Name()]
		if !ok {
			log.Printf("topic not found: %s", e.Name())
			continue
		}

		body, err := e.ToBody()
		if err != nil {
			log.Printf("failed to marshal event %v", err)
			continue
		}
		if err := s.sns.PublishMessage(ctx, topic, body); err != nil {
			// TODO Transactional Outbox Patternの導入までは通知して手動対応ってなるのだろうか。。。
			log.Printf("failed to publish message %v", err)
		}
	}
}
