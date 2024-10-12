package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	// DB
	DSN string `env:"KITCHEN_API_DSN"`

	// SNS
	ArnTopicEventKitchenTicketCreated        string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_CREATED"`
	ArnTopicEventKitchenTicketCreationFailed string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_CREATION_FAILED"`
	ArnTopicEventKitchenTicketApproved       string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_APPROVED"`
	ArnTopicEventKitchenTicketRejected       string `env:"TOPIC_ARN_EVENT_KITCHEN_TICKET_REJECTED"`

	// SQS
	SqsUrlTicketEvent   string `env:"SQS_URL_TICKET_EVENT"`
	SqsUrlTicketCommand string `env:"SQS_URL_TICKET_COMMAND"`
	SqsUrlTicketReply   string `env:"SQS_URL_TICKET_REPLY"`
}

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
