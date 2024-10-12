package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	// DB
	DSN string `env:"KITCHEN_API_DSN"`

	// SNS
	ArnTopicCommandKitchenTicketCreate  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_CREATE"`
	ArnTopicCommandKitchenTicketApprove string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_APPROVE"`
	ArnTopicCommandKitchenTicketReject  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_REJECT"`

	ArnTopicEventKitchenTicketCreated        string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_CREATED"`
	ArnTopicEventKitchenTicketCreationFailed string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_CREATION_FAILED"`
	ArnTopicEventKitchenTicketApproved       string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_APPROVED"`
	ArnTopicEventKitchenTicketRejected       string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_REJECTED"`

	// SQS
	SqsUrlOrderReply    string `env:"SQS_URL_ORDER_REPLY"`
	SqsUrlTicketCommand string `env:"SQS_URL_TICKET_COMMAND"`
}

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
