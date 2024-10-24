package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	// DB
	DSN string `env:"BILLING_API_DSN"`

	// SNS
	ArnTopicEventBillingCardAuthorized      string `env:"TOPIC_ARN_EVENT_BILLING_CARD_AUTHORIZED"`
	ArnTopicEventBillingCardAuthorizeFailed string `env:"TOPIC_ARN_EVENT_BILLING_CARD_AUTHORIZE_FAILED"`
	ArnTopicEventBillingCardCanceled        string `env:"TOPIC_ARN_EVENT_BILLING_CARD_CANCELED"`

	// SQS
	SqsUrlBillingEvent   string `env:"SQS_URL_BILLING_EVENT"`
	SqsUrlBillingCommand string `env:"SQS_URL_BILLING_COMMAND"`
	SqsUrlBillingReply   string `env:"SQS_URL_BILLING_REPLY"`
}

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
