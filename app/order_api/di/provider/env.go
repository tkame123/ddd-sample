package provider

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type ENV = string

const (
	EnvDevelopment ENV = "development"
	EnvProduction  ENV = "production"
)

type EnvConfig struct {
	ENV string `env:"ENV" envDefault:"production"`

	// DB
	DbDebug bool   `env:"DB_DEBUG" envDefault:"false"`
	DSN     string `env:"ORDER_API_DSN"`

	// AUTH0
	AuthDomainName   string `env:"AUTH0_DOMAIN"`
	AuthAudienceName string `env:"AUTH0_AUDIENCE_ORDER_API"`

	// SNS
	ArnTopicEventOrderOrderCreated               string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CREATED"`
	ArnTopicEventOrderOrderApproved              string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_APPROVED"`
	ArnTopicEventOrderOrderRejected              string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_REJECTED"`
	ArnTopicEventOrderOrderCanceled              string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CANCELED"`
	ArnTopicEventOrderOrderCancellationConfirmed string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CANCELLATION_CONFIRMED"`
	ArnTopicEventOrderOrderCancellationRejected  string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CANCELLATION_REJECTED"`

	ArnTopicCommandOrderOrderApprove    string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_APPROVE"`
	ArnTopicCommandOrderOrderReject     string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_REJECT"`
	ArnTopicCommandKitchenTicketCreate  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_CREATE"`
	ArnTopicCommandKitchenTicketApprove string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_APPROVE"`
	ArnTopicCommandKitchenTicketReject  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_REJECT"`
	ArnTopicCommandKitchenTicketCancel  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_CANCEL"`

	ArnTopicCommandBillingCardAuthorize string `env:"TOPIC_ARN_COMMAND_BILLING_CARD_AUTHORIZE"`
	ArnTopicCommandBillingCardCancel    string `env:"TOPIC_ARN_COMMAND_BILLING_CARD_CANCEL"`

	// SQS
	SqsUrlOrderEvent   string `env:"SQS_URL_ORDER_EVENT"`
	SqsUrlOrderCommand string `env:"SQS_URL_ORDER_COMMAND"`
	SqsUrlOrderReply   string `env:"SQS_URL_ORDER_REPLY"`
}

func NewENV() (*EnvConfig, error) {
	var cfg EnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &cfg, nil
}
