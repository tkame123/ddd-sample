package provider

import "github.com/tkame123/ddd-sample/proto/message"

type TopicArn = string

type PublisherConfig struct {
	TopicMap map[message.Type]TopicArn
}

func NewPublisherConfig(envCfg *EnvConfig) *PublisherConfig {
	return &PublisherConfig{
		TopicMap: map[message.Type]TopicArn{
			message.Type_TYPE_EVENT_CARD_AUTHORIZED:           envCfg.ArnTopicEventBillingCardAuthorized,
			message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED: envCfg.ArnTopicEventBillingCardAuthorizeFailed,
		},
	}
}
