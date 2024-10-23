package provider

import "github.com/tkame123/ddd-sample/proto/message"

type TopicArn = string

type PublisherConfig struct {
	TopicMap map[message.Type]TopicArn
}

func NewPublisherConfig(envCfg *EnvConfig) *PublisherConfig {
	return &PublisherConfig{
		TopicMap: map[message.Type]TopicArn{
			message.Type_TYPE_EVENT_TICKET_CREATED:               envCfg.ArnTopicEventKitchenTicketCreated,
			message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:       envCfg.ArnTopicEventKitchenTicketCreationFailed,
			message.Type_TYPE_EVENT_TICKET_APPROVED:              envCfg.ArnTopicEventKitchenTicketApproved,
			message.Type_TYPE_EVENT_TICKET_REJECTED:              envCfg.ArnTopicEventKitchenTicketRejected,
			message.Type_TYPE_EVENT_TICKET_CANCELED:              envCfg.ArnTopicEventKitchenTicketCanceled,
			message.Type_TYPE_EVENT_TICKET_CANCELLATION_REJECTED: envCfg.ArnTopicEventKitchenTicketCancellationRejected,
		},
	}
}
