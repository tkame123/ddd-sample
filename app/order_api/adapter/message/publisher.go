package message

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event"
	"log"
)

type topicArn = string

type arnConfig struct {
	ArnTopicEventOrderOrderCreated      string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CREATED"`
	ArnTopicEventOrderOrderApproved     string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_APPROVED"`
	ArnTopicEventOrderOrderRejected     string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_REJECTED"`
	ArnTopicCommandOrderOrderApprove    string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_APPROVE"`
	ArnTopicCommandOrderOrderReject     string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_REJECT"`
	ArnTopicCommandKitchenTicketCreate  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_CREATE"`
	ArnTopicCommandKitchenTicketApprove string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_APPROVE"`
	ArnTopicCommandKitchenTicketReject  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_REJECT"`
	ArnTopicCommandBillingCardAuthorize string `env:"TOPIC_ARN_COMMAND_BILLING_CARD_AUTHORIZE"`
}

type eventPublisher struct {
	sns      sns.Actions
	topicMap map[event.Name]topicArn
}

func NewEventPublisher(sns *sns.Actions) domain_event.Publisher {
	var arnCfg arnConfig
	if err := env.Parse(&arnCfg); err != nil {
		log.Fatalf("unable to parase env, %v", err)
	}

	topicMap := map[event.Name]topicArn{
		event.EventName_OrderCreated:    arnCfg.ArnTopicEventOrderOrderCreated,
		event.EventName_OrderApproved:   arnCfg.ArnTopicEventOrderOrderApproved,
		event.EventName_OrderRejected:   arnCfg.ArnTopicEventOrderOrderRejected,
		event.CommandName_TicketCreate:  arnCfg.ArnTopicCommandKitchenTicketCreate,
		event.CommandName_TicketApprove: arnCfg.ArnTopicCommandOrderOrderApprove,
		event.CommandName_TicketReject:  arnCfg.ArnTopicCommandOrderOrderReject,
		event.CommandName_CardAuthorize: arnCfg.ArnTopicCommandBillingCardAuthorize,
	}

	return &eventPublisher{sns: *sns, topicMap: topicMap}
}

func (s *eventPublisher) PublishMessages(ctx context.Context, events []event.Event) {
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
