package message

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event"
	"log"
)

var topicMap = map[event.Name]sns.Topic{
	event.EventName_OrderCreated:   sns.Topic_Event_Order_OrderCreated,
	event.EventName_OrderApproved:  sns.Topic_Event_Order_OrderrApproved,
	event.EventName_OrderRejected:  sns.Topic_Event_Order_OrderRejected,
	event.EventName_TicketCreated:  sns.Topic_Command_Kitchen_TicketCreate,
	event.EventName_TicketApproved: sns.Topic_Command_Kitchen_TicketAprove,
	event.EventName_TicketRejected: sns.Topic_Command_Kitchen_TicketReject,
	event.EventName_CardAuthorized: sns.Topic_Command_Billing_CardAuthorize,
}

type eventPublisher struct {
	sns sns.Actions
}

func NewEventPublisher(sns *sns.Actions) domain_event.Publisher {
	return &eventPublisher{sns: *sns}
}

func (s *eventPublisher) PublishMessages(ctx context.Context, events []model.OrderEvent) {
	for _, e := range events {
		var topic sns.Topic
		topic, ok := topicMap[e.Name()]
		if !ok {
			log.Printf("topic not found: %s", e.Name())
			continue
		}
		// TODO: Bodyの作成
		if err := s.sns.PublishMessage(ctx, topic, e.Name()); err != nil {
			// TODO Transactional Outbox Patternの導入までは通知して手動対応ってなるのだろうか。。。
			log.Printf("failed to publish message %v", err)
		}
	}
}