package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	OrderAPIDSN string `env:"ORDER_API_DSN"`

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

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
