package provider

import "github.com/tkame123/ddd-sample/proto/message"

type TopicArn = string

type PublisherConfig struct {
	TopicMap map[message.Type]TopicArn
}

func NewPublisherConfig(envCfg *EnvConfig) *PublisherConfig {
	return &PublisherConfig{
		TopicMap: map[message.Type]TopicArn{
			message.Type_TYPE_EVENT_ORDER_CREATED:    envCfg.ArnTopicEventOrderOrderCreated,
			message.Type_TYPE_EVENT_ORDER_APPROVED:   envCfg.ArnTopicEventOrderOrderApproved,
			message.Type_TYPE_EVENT_ORDER_REJECTED:   envCfg.ArnTopicEventOrderOrderRejected,
			message.Type_TYPE_COMMAND_TICKET_CREATE:  envCfg.ArnTopicCommandKitchenTicketCreate,
			message.Type_TYPE_COMMAND_TICKET_APPROVE: envCfg.ArnTopicCommandKitchenTicketApprove,
			message.Type_TYPE_COMMAND_TICKET_REJECT:  envCfg.ArnTopicCommandKitchenTicketReject,
			message.Type_TYPE_COMMAND_CARD_AUTHORIZE: envCfg.ArnTopicCommandBillingCardAuthorize,
		},
	}
}
